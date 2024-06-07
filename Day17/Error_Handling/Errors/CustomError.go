package errors

import "fmt"

type CustomError struct {
	ErrMessage string
	Code       int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error message :%v\nError Code :%v", e.ErrMessage, e.Code)
}
