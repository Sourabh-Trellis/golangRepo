package main

import "fmt"

func main()  {
	var i,j string = "hello","sourabh"
	var l,m int = 10,20
	fmt.Print(i)
	fmt.Print(j)    //prints args om same line without spaces between them

	fmt.Print(i,j)  //prints args om same line without spaces between them

	fmt.Print(l,m)  //is args are not string it automatically adds space between them
	fmt.Println()
	fmt.Println(i,j)  //prints white space between args and newline at the end
	fmt.Println(l,m)	//prints white space between args and newline at the end

	fmt.Printf("i has value:%v and type:%T",i,i)	//formats the args based on formatting 
													//verb and then prints them


}