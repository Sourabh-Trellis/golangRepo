package main

import (
	errors "Errors/Errors"
	functions "Errors/Functions"
	"fmt"
)

func main() {

	var num int
	fmt.Println("Enter number you want to divide:")
	fmt.Scan(&num)

	var divisor int
	fmt.Println("Enter the divisor:")
	fmt.Scan(&divisor)

	result, err := functions.Divide(float32(num), float32(divisor))
	if err != nil {
		var divideErr errors.DivideByZeroError
		divideErr.ErrMessage = err.Error()
		divideErr.Code = 302
		fmt.Println(divideErr.Error())
	} else {
		fmt.Println(result)
	}
}
