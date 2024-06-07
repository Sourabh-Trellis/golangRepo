package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	Sourabh := User{Name: "sourabh", Age: 25, Email: "sourabh@gmail.com", Status: true}
	fmt.Println(Sourabh)

	fmt.Printf("%+v \n",Sourabh)

	fmt.Println(Sourabh.Name,Sourabh.Email,Sourabh.Age,Sourabh.Status)

	abc()
}

func abc () {
	fmt.Println("hi")
}
