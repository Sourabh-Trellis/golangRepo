package main

import "fmt"

// func message(msg string) {

// 	fmt.Println(msg)
// }

// func addition(a int,b int) (result int) {
// 	result = a+b
// 	return result
// }

// func myFunction(a string, b int) (product int, txt string) {
// 	product = b * b
// 	txt = "Hello " + a
// 	return
// }

func factorial(x int) (y int) {
	
	if( x > 0 ){
		y = x * factorial(x-1)
	} else {
		y = 1
	}

	return
}

func main() {

	// fmt.Println("Enter  a text :")
	// var msg string
	// fmt.Scan(&msg)
	// message(msg)

	// fmt.Println(addition(5,10))

	// a,b := myFunction("sourabh",5)

	// fmt.Println(a,b)


	fmt.Println("enter a number to get factorial : ")
	var num int
	fmt.Scan(&num)
	fmt.Println(factorial(num))
}
