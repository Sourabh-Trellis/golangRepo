package main

import "fmt"

func main() {

	myslice := []int{10, 20, 30}	//slice declaration and initialization
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	

	// var arr1 = [5]int {1,2,3,4,5}	//creating slice from array
	// myslice1 := arr1[2:]				//arr1[start index,end index]
	// fmt.Println(myslice1)
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))


	myslice2 := make([]int,5,10)	//creating array using make(type,length,capacity) function
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

	myslice2[2]=20		//changing the elements of slice 
	myslice2[4]=80
	fmt.Println(myslice2)

	fmt.Println(myslice2[2])	//accessing elements of slice
	fmt.Println(myslice2[3])

	myslice2 = append(myslice2, 60)		//append an element to slice
	fmt.Println(myslice2)
	
	myslice3 := append(myslice,myslice2...)
	fmt.Println(myslice3)

	myslice3 = append(myslice3, 22)
	fmt.Println(myslice3)

	
}
