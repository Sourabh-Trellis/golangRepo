package main

import (
	"container/list"
	"fmt"
)

func main() {

	myList := list.New()

	fmt.Println(myList.Back())
	fmt.Println(myList.Front())

	//Adding elements at back of list
	ele1 := myList.PushBack(2)
	ele2 := myList.PushBack("ddd")
	myList.PushBack(22)
	myList.PushBack("hfhfhf")
	//Adding elements at front of list
	myList.PushFront(true)
	myList.PushFront("sourabh")

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//Removing element from front of list
	if firstElement := myList.Front(); firstElement != nil {
		myList.Remove(firstElement)
	}
	fmt.Println("----")

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//Removing element from back of the list
	if lastElement := myList.Back(); lastElement != nil {
		myList.Remove(lastElement)
	}

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//returns value stored in first element
	fmt.Println(myList.Front().Value)
	//returns first element ,nil if not present
	fmt.Println(myList.Front())

	//returns value stored in last element
	fmt.Println(myList.Back().Value)
	//returns last element ,nil if not present
	fmt.Println(myList.Back())

	//insert after particular element
	fmt.Println(myList.InsertAfter("Naruto", ele1))

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//insert before particular element
	fmt.Println(myList.InsertBefore("itachi", ele2))

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//move element to end of list
	myList.MoveToBack(ele2)

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//move element to back of list
	myList.MoveToFront(ele1)

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//insert copy of list(same or other) at end
	myList.PushBackList(myList)

	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//insert copy of list(same or other) at front
	myList.PushFrontList(myList)
	//iterate over the list and print elements
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("----")

	//remove element from list
	fmt.Println(myList.Remove(ele1))
	
}
