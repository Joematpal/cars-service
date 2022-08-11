package carsdb

import (
	"context"

	"github.com/joematpal/cars-service/internal/logging"
	cars "github.com/joematpal/cars-service/pkg/cars/v1"
	sqlp "github.com/joematpal/go-sql/v2"
	"github.com/scylladb/gocqlx/qb"
)

type DB struct {
	db       *sqlp.DB
	debugger logging.Debugger
}

func New(opts ...DBOption) (*DB, error) {
	db := &DB{}
	for _, opt := range opts {
		if err := opt.applyOption(db); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func (db *DB) Debugf(format string, args ...interface{}) {
	if db.debugger != nil {
		db.debugger.Debugf(format, args...)
	}
}

func (db *DB) CreateCar(ctx context.Context, req *cars.Car) (*cars.Car, error) {
	// TODO: add update/create times
	// TODO: add debugger

	stmt, names := qb.Insert(CarsTable.Name).Columns(CarsTable.ListColumns()...).ToCql()

	if err := db.db.Exec(stmt, names, req); err != nil {
		return nil, err
	}
	return req, nil
}

func (db *DB) ListCars(ctx context.Context, req *cars.ListCarsReq) (*cars.Cars, error) {
	stmt, names := qb.Select(CarsTable.Name).Columns(CarsTable.ListColumns()...).ToCql()

	list := []*cars.Car{}
	if err := db.db.Select(&list, stmt, names, req); err != nil {
		return nil, err
	}

	return &cars.Cars{
		Cars: list,
	}, nil
}

func (db *DB) GetCar(ctx context.Context, req *cars.CarReq) (*cars.Car, error) {
	stmt, names := qb.Select(CarsTable.Name).Columns(CarsTable.ListColumns()...).ToCql()

	out := &cars.Car{}
	if err := db.db.Get(out, stmt, names, req); err != nil {
		return nil, err
	}

	return out, nil
}

func (db *DB) DeleteCar(ctx context.Context, req *cars.CarReq) error {
	stmt, names := qb.Delete(CarsTable.Name).Where(qb.Eq(ColumnKey_ID)).ToCql()

	if err := db.db.Exec(stmt, names, req); err != nil {
		return err
	}
	return nil
}
