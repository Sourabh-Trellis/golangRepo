package functions

import "errors"

func Divide(num float32, divisor float32) (float32, error) {

	if divisor != 0 {
		return num / divisor, nil
	} else {
		return 0.0, errors.New("cannot divide by zero")
	}
}
