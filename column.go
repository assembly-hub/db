package db

import "reflect"

type ColumnType interface {
	Name() string
	Nullable() (nullable, ok bool)
	ScanType() reflect.Type
	DatabaseTypeName() string
	DecimalSize() (precision, scale int64, ok bool)
	Length() (length int64, ok bool)
}
