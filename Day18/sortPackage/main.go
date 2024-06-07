package main

import (
	"fmt"
	"sort"
)

func main() {

	//defining a slice
	nums := []int{5, 7, 2, 8, 4, 9}

	//returns true if slice is sorted else false
	fmt.Println(sort.IntsAreSorted(nums))

	//sort the integer slice in ascending order
	sort.Ints(nums)

	//printing sorted slice
	fmt.Println("sorted slice :", nums)
	//returns true if slice is sorted else false
	fmt.Println(sort.IntsAreSorted(nums))

	//defining a slice of string
	words := []string{"banana", "apple", "cherry", "date"}

	//sorting the slice
	sort.Strings(words)

	//printing sorted slice
	fmt.Println("sorted slice :", words)

}
