package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hi")
	defer fmt.Println("how are you")
	defer fmt.Println("how you doin")
	calc(2, 3)
}
func calc(a int, b int) {
	result := a + b
	fmt.Println(result)
}
