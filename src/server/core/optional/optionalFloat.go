package optional

import (
	"database/sql/driver"
	"fmt"
)

type Float32 float32
type Float64 float64

func (dst Float32) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case float64:
		return Float32(src), nil
	case float32:
		return Float32(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Float32) Value() (driver.Value, error) {
	return float32(src), nil
}

func (dst Float64) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case float64:
		return Float64(src), nil
	case float32:
		return Float64(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Float64) Value() (driver.Value, error) {
	return float64(src), nil
}
