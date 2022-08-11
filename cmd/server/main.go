package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	carsdb "github.com/joematpal/cars-service/internal/db/v1/cars"
	carspb "github.com/joematpal/cars-service/pkg/cars/v1"

	loggerf "github.com/joematpal/go-logger/flags"
	grpcf "github.com/joematpal/go-server/flags"
	grpcp "github.com/joematpal/go-server/grpc"
	sqlp "github.com/joematpal/go-sql/v2"
	sqlf "github.com/joematpal/go-sql/v2/flags"

	cli "github.com/urfave/cli/v2"

	"github.com/joematpal/cars-service/internal/cars/v1"
	"github.com/joematpal/cars-service/internal/flags"
	"github.com/joematpal/cars-service/internal/version"

	"github.com/joematpal/go-logger"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
	}()
	if err := NewApp().RunContext(ctx, os.Args); err != nil {
		panic(err)
	}
}

// NewApp wrapper for the entrypoint of the app
func NewApp() *cli.App {
	return &cli.App{
		Name: "cars-service",
		Commands: []*cli.Command{
			versionCMD,
			{
				Name:  "server",
				Usage: "starts the cars-service api service",
				Flags: flags.Join(
					sqlf.DBFlags,
					grpcf.GRPCFlags,
					loggerf.LogFlags,
				),
				Action: func(c *cli.Context) error {
					logr, err := setupLogger(c)
					if err != nil {
						return fmt.Errorf("setup logger: %v", err)
					}

					logr.Debug("Connect to accounts and subscriptions db's")
					db, err := sqlp.New(
						sqlp.WithType(c.String(sqlf.DBType)),
						sqlp.WithDBName(c.String(sqlf.DBName)),
						sqlp.WithHosts(c.String(sqlf.DBHosts)),
						sqlp.WithUser(c.String(sqlf.DBUser)),
						sqlp.WithPassword(c.String(sqlf.DBPass)),
						sqlp.WithPort(c.String(sqlf.DBPort)),
					)
					if err != nil {
						return err
					}

					carsDB, err := carsdb.New(
						carsdb.WithDB(db),
						carsdb.WithDebugger(logr),
					)
					if err != nil {
						return fmt.Errorf("new car db: %v", err)
					}
					if err != nil {
						logr.Fatal(err)
					}

					carsServer, err := cars.NewServer(
						cars.WithLogger(logr),
						cars.WithDB(carsDB),
					)
					if err != nil {
						return fmt.Errorf("new car server: %v", err)
					}

					grpcOpts := setupBaseGRPC(c, logr,
						func(s *grpc.Server) {
							carspb.RegisterCarsServiceServer(s, carsServer)
						},
						[]grpcp.GatewayServiceHandler{
							carspb.RegisterCarsServiceHandler,
						},

						[]grpc.StreamServerInterceptor{},
						[]grpc.UnaryServerInterceptor{},
					)
					srv, err := grpcp.New(grpcOpts...)
					if err != nil {
						return fmt.Errorf("new grpcp: %v", err)
					}

					logr.Infof("Server started")
					if err := srv.StartWithContext(c.Context); err != nil {
						return fmt.Errorf("start: %v", err)
					}

					return nil
				},
			},
		},
		Action: func(app *cli.Context) error {
			if app.Bool("version") {
				fmt.Print(version.Version)
				return nil
			}

			cli.ShowAppHelpAndExit(app, 0)
			return nil
		},
	}
}

func setupLogger(c *cli.Context) (logger.CorrelationLogger, error) {
	// Setup Logger
	logOpts := []logger.Option{
		logger.WithEnv(c.String(loggerf.LogEnv)),
		logger.WithLevel(c.String(loggerf.LogLevel)),
		logger.WithLogStacktrace(c.Bool(loggerf.LogStacktrace)),
	}
	logr, err := logger.New(logOpts...)
	if err != nil {
		return nil, fmt.Errorf("new logger: %v", err)
	}
	return logr, nil
}

func setupBaseGRPC(
	c *cli.Context,
	logr logger.CorrelationLogger,
	registerServiceFunc func(*grpc.Server),
	gatewayServiceHandlers []grpcp.GatewayServiceHandler,
	streamInterceptors []grpc.StreamServerInterceptor,
	unaryInterceptors []grpc.UnaryServerInterceptor,
) []grpcp.Option {
	dialOpts := []grpc.DialOption{}
	pubCert := c.String(grpcf.GRPCPubCert)
	privCert := c.String(grpcf.GRPCPrivCert)

	grpcGWHost := c.String(grpcf.GRPCGatewayClientHost)
	grpcGWPort := c.String(grpcf.GRPCGatewayClientPort)

	opts := []grpcp.Option{
		grpcp.WithTLS(c.Bool(grpcf.GRPCTLS)),
		grpcp.WithServerOptions(
			grpc.ChainStreamInterceptor(
				append([]grpc.StreamServerInterceptor{
					logger.LoggingStreamServerInterceptor(logr),
				}, streamInterceptors...)...,
			),
			grpc.ChainUnaryInterceptor(append([]grpc.UnaryServerInterceptor{
				logger.LoggingUnaryServerInterceptor(logr),
			}, unaryInterceptors...)...),
		),
		grpcp.WithRegisterService(registerServiceFunc),
		grpcp.WithGatewayServiceHandlers(
			gatewayServiceHandlers...,
		),
		grpcp.WithLogger(logr),
		grpcp.WithGatewayServerMuxOptions(),
	}

	swaggerFile := c.String(grpcf.SwaggerFile)

	if swaggerFile != "" {
		swaggerFile = "swagger/cars.swagger.json"
	}

	opts = append(opts, grpcp.WithSwaggerFile(swaggerFile))

	if grpcGWHost != "" && grpcGWPort != "" {
		opts = append(opts, grpcp.WithGatewayAddr(grpcGWHost, grpcGWPort))
	}

	if privCert == "" && pubCert == "" {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}

	// NOTE: If we pass more than one option in the gateway dial options we need to also pass in the tls stuff
	opts = append(opts, grpcp.WithGatewayDialOptions(
		dialOpts...,
	))

	if host := c.String(grpcf.GRPCHost); host != "" {
		opts = append(opts, grpcp.WithHost(host))
	}

	if port := c.String(grpcf.GRPCPort); port != "" {
		opts = append(opts, grpcp.WithPort(port))
	}

	if tls := c.Bool(grpcf.GRPCTLS); tls {
		opts = append(opts, grpcp.WithTLS(tls))
	}

	if pubCert != "" {
		opts = append(opts, grpcp.WithPubCert(pubCert))
	}

	if privCert != "" {
		opts = append(opts, grpcp.WithPrivCert(privCert))
	}

	return opts
}

var versionCMD = &cli.Command{
	Name: "version",
	Action: func(app *cli.Context) error {
		fmt.Printf(
			"Version: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n",
			version.Version,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		)
		return nil
	},
}
