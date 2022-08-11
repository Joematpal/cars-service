package cars

import (
	"context"
	"net/http"

	"github.com/joematpal/cars-service/internal/logging"
	cars "github.com/joematpal/cars-service/pkg/cars/v1"
)

type FactoryDB interface {
	CreateCar(ctx context.Context, req *cars.Car) (*cars.Car, error)
	ListCars(ctx context.Context, req *cars.ListCarsReq) (*cars.Cars, error)
	GetCar(ctx context.Context, req *cars.CarReq) (*cars.Car, error)
	DeleteCar(ctx context.Context, req *cars.CarReq) error
}

type Server struct {
	// do something with the logger..
	logger logging.Logger
	fdb    FactoryDB
	cars.UnimplementedCarsServiceServer
}

func NewServer(opts ...Option) (*Server, error) {
	cp := &Server{}

	for _, o := range opts {
		if err := o.applyOption(cp); err != nil {
			return nil, err
		}
	}
	return cp, nil
}

func (s *Server) CreateCar(ctx context.Context, req *cars.Car) (*cars.Car, error) {

	return s.fdb.CreateCar(ctx, req)
}

func (s *Server) ListCars(ctx context.Context, req *cars.ListCarsReq) (*cars.Cars, error) {
	return s.fdb.ListCars(ctx, req)

}

func (s *Server) GetCar(ctx context.Context, req *cars.CarReq) (*cars.Car, error) {
	return s.fdb.GetCar(ctx, req)
}

func (s *Server) DeleteCar(ctx context.Context, req *cars.CarReq) (*cars.Success, error) {
	if err := s.fdb.DeleteCar(ctx, req); err != nil {
		return nil, err
	}

	return &cars.Success{
		Status:  http.StatusAccepted,
		Message: "deleted",
	}, nil
}
