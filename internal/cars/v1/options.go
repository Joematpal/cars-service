package cars

import "github.com/joematpal/cars-service/internal/logging"

type Option interface {
	applyOption(*Server) error
}

type applyOptionFunc func(*Server) error

func (f applyOptionFunc) applyOption(s *Server) error {
	return f(s)
}

func WithDB(fdb FactoryDB) Option {
	return applyOptionFunc(func(s *Server) error {
		s.fdb = fdb
		return nil
	})
}

// WithLogger sets the logger
func WithLogger(logger logging.Logger) Option {
	return applyOptionFunc(func(cp *Server) error {
		cp.logger = logger
		return nil
	})
}
