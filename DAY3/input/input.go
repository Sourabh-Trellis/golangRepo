package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome sourabh"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza")

	//comma ok || err ok
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating",input)
}
