package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var number int64
	// var input string
	// fmt.Println("enter a number-")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("err reading input ", err)
	// 	return
	// }

	// input = strings.TrimSpace(input)
	// _, err = strconv.ParseInt(input, 36, 64)
	// if err != nil {
	// 	fmt.Println("invalid input,plz enter a number")
	// } else {
	// 	number, _ = strconv.ParseInt(input, 36, 64)
	// }

	// if number < 0 {
	// 	fmt.Println("number you entered is negative.")
	// } else if number > 0 {
	// 	fmt.Println("number you entered is positive.")
	// } else {
	// 	fmt.Println("number entered is zero")
	// }

	var weekday int64

	var input string
	fmt.Println("plz enter number 1-7 -")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	input = strings.TrimSpace(input)

	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Invalid input.", err)
		return
	} else {
		weekday, _ = strconv.ParseInt(input, 10, 64)
		if weekday > 7 {
			fmt.Println(weekday)
			fmt.Println("Invalid number")
		} else {
			switch weekday {
			case 1:
				fmt.Println("Sunday")
			case 2:
				fmt.Println("Monday")
			case 3:
				fmt.Println("Tuesday")
			case 4:
				fmt.Println("Wednesday")
			case 5:
				fmt.Println("Thursday")
			case 6:
				fmt.Println("Friday")
			case 7:
				fmt.Println("Saturday")
			}
		}
	}
}
