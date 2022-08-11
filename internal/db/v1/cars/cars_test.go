package carsdb

import (
	"context"
	"reflect"
	"testing"

	cars "github.com/joematpal/cars-service/pkg/cars/v1"
)

func TestDB_CreateCar(t *testing.T) {

	type args struct {
		ctx context.Context
		req *cars.Car
	}
	tests := []struct {
		name    string
		args    args
		db      *DB
		want    *cars.Car
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.db.CreateCar(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.CreateCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.CreateCar() = %v, want %v", got, tt.want)
			}
		})
	}
}
