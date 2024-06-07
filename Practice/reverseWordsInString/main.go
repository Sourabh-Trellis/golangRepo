package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your string:")

	input, _ := reader.ReadString('\n')

	reversed := reverseWordsInString(input)

	fmt.Println("Reversed string:")
	fmt.Println(reversed)

}

func reverseWordsInString(in string) string {

	split := strings.Split(in, " ")

	for idx, val := range split {

		runes := []rune(val)
		result := make([]rune, len(runes))
		j := 0

		for i := len(runes) - 1; i >= 0; i-- {
			result[j] = runes[i]
			j++
		}
		split[idx] = string(result)
	}
	return strings.Join(split, " ")
}
