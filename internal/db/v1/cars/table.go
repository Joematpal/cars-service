package carsdb

import (
	table "github.com/joematpal/go-sql/v2/table"
)

type ColumnKey = string

const (
	ColumnKey_ID ColumnKey = "id"
)

var CarsTable = table.New("cars", map[string]struct{}{
	ColumnKey_ID: {},
})
