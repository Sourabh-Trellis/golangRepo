package main

import (
	"fmt"
	"strconv"
)

func main() {

	//defining string containing numbers
	numstr := "12345"

	//converting above string in integer
	num, err := strconv.ParseInt(numstr, 10, 64)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Printf("Parsed Int :%T\n", num)

	//defining string containing float
	floatstr := "3.14"

	//converting above string to float
	float, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Printf("parsed float :%v %T\n", float, float)

	//converting int to string
	str := strconv.Itoa(1234)
	fmt.Println("converted to string:", str)
}
