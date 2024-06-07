package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza shop")
	fmt.Println("please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating, ", rating)

	numrating , err := strconv.ParseFloat(strings.TrimSpace(rating),64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating, ",numrating + 1)

	}

}
