package main

import (
	"fmt"
	"os"
)

func main() {
	name, age := "sourabh", 29
	bytes, err := fmt.Fprint(os.Stdout, name, age)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Fprint err :%v", err)
	}
	fmt.Println("bytes printed :", bytes)
}
