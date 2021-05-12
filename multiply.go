package multiply

import (
	"errors"
	"math"
)

func validateNumber(c *float32) error {
	// nil prevention validation
	if c == nil {
		return errors.New("Number is nil")
	}
	// check number is inf
	if math.IsInf(float64(*c), 0) {
		return errors.New("Number is really big and exceeded maximum range value")
	}
	return nil
}

func Multiply(a *float32, b *float32) (*float32, error) {
	var c float32 = (*a) * (*b)
	err := validateNumber(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
