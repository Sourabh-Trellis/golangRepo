package main

import "fmt"

func main()  {
	
//array declaration

// 1) with var keyword

	// var arr1 =[5]int{1,2,3,4,5}
	// var arr2 =[...]string{"pune","mumbai"}

	// fmt.Println(arr1)
	// fmt.Println(arr2)

// 2) with := sign

	// arr3 := [5]string{"pune","mumbai","gujrat","gandhinagar","infocity"}
	// arr4 := [...]int{11,22,33,44}

	// fmt.Println(arr3)
	// fmt.Println(arr4)
	
//array initialization

// Not initialized

	// arr5 := [5]int{}	//initializes with default values
	// fmt.Println(arr5)

// Partially initialized

	// arr6 := [5]string{"sourabh","malhar"}
	// fmt.Println(arr6)

	// arr7 := [5]int{1,2}
	// fmt.Println(arr7)

// fully iniyialized 

	arr8 := [5]int{1,2,3,4,5}
	// fmt.Println(arr8)

// accessing the elements in an array

	// fmt.Println(arr8[2])
	// fmt.Println(arr8[6])

// changing the elements in an array

	// arr8[3]=100
	// fmt.Println(arr8)
	// arr8[0]=200
	// fmt.Println(arr8)

// initialize only specific elements

	array := [8]int{1:20,5:20}
	fmt.Println(array)

	fmt.Println(len(arr8))
	fmt.Println(len(array))

}