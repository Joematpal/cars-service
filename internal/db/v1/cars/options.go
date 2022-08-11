package carsdb

import (
	"github.com/joematpal/cars-service/internal/logging"
	sqlp "github.com/joematpal/go-sql/v2"
)

// DBOption interface to add values to the db struct
type DBOption interface {
	applyOption(*DB) error
}

type applyOptionFunc func(*DB) error

func (f applyOptionFunc) applyOption(db *DB) error {
	return f(db)
}

// WithDatastore sets the data store
func WithDB(dbr *sqlp.DB) DBOption {
	return applyOptionFunc(func(d *DB) error {
		d.db = dbr
		return nil
	})
}

// WithDebugger sets the logger
func WithDebugger(l logging.Debugger) DBOption {
	return applyOptionFunc(func(d *DB) error {
		d.debugger = l
		return nil
	})
}
