package main

import "fmt"

func main() {

	var a = map[int]string{1: "sourabh", 2: "Sangam", 3: "kirti"}

	fmt.Println(a)

	b := make(map[int]string)

	b[1] = "abcd"
	b[2] = "efgh"
	b[3] = "ijkl"

	fmt.Println(b)

	a[2] = "apeksha"
	b[2] = "ddddddddddddddddd"

	// fmt.Println(a)
	// fmt.Println(b)

	// fmt.Println(a[3])

	val, ok := a[4]
	fmt.Println(val, ok)


}
