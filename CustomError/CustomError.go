package customerror

import "fmt"

type CustomError struct {
	status  bool
	message string
	code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Status:%v\nMessage:%v\nCode:%v", e.status, e.message, e.code)
}

