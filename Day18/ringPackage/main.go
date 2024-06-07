package main

import (
	"container/ring"
	"fmt"
)

func main() {

	//create a new ring with particular number of elements
	myRing := ring.New(4)

	//initialize ring with some values
	for i := 1; i <= myRing.Len(); i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	//iterate over the ring and print the elements
	fmt.Println("rinig elements")
	myRing.Do(func(x interface{}) {
		fmt.Println(x)
	})

	//move ring to next element
	myRing = myRing.Move(2)

	//iterate over ring and print elements
	fmt.Println("after moving twice")
	myRing.Do(func(x interface{}) {
		fmt.Println(x)
	})

	myRing = myRing.Move(1)
	fmt.Println("After moving once")
	myRing.Do(func (x interface{})  {
		fmt.Println(x)
	})

}
