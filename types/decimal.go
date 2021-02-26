package types

import (
	"database/sql/driver"
	"errors"
	"github.com/shopspring/decimal"
)

const Precision = 10e8

type EntityDecimal struct {
	d decimal.Decimal
}

func NewEntityDecimalFromDecimal(d decimal.Decimal) EntityDecimal {
	return EntityDecimal{d: d}
}

func NewEntityDecimalFromInt64(i int64) EntityDecimal {
	return NewEntityDecimalFromDecimal(decimal.NewFromInt(i))
}

func NewEntityDecimalFromFloat(f float64) EntityDecimal {
	return NewEntityDecimalFromDecimal(decimal.NewFromFloat(f))
}

func (e *EntityDecimal) Decimal() decimal.Decimal {
	return e.d
}

func (e EntityDecimal) Value() (driver.Value, error) {
	return e.d.Mul(decimal.New(Precision, 0)).IntPart(), nil
	//f, _ := e.d.Float64()
	//return int64(f * Precision), nil
}

func (e *EntityDecimal) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	if i, ok := src.(int64); ok {
		e.d = decimal.New(i, 0).Div(decimal.New(Precision, 0))
		//e.d = decimal.NewFromFloat(float64(i) / Precision)
		return nil
	} else {
		return errors.New("only support int64")
	}
}
