package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your string:")
	input, _ := reader.ReadString('\n')

	reversedString := reverseString(input)
	fmt.Println("Reverserd string:")
	fmt.Println(reversedString)
}

func reverseString(in string) string {
	runes := []rune(in)

	result := make([]rune, len(runes))
	j := 0

	for i := len(runes) - 1; i >= 0; i-- {
		result[j] = runes[i]
		j++
	}
	return string(result)
}
