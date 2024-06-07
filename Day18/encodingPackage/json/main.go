package main

import (
	"encoding/json"
	"fmt"
)

//define struct
// json:"name" allows you to customize the corresponding field name in the encoded JSON.

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {

	//create a person struct
	person := Person{Name: "Sourabh Deshmukh", Age: 25, City: "Pune"}

	//Encode the struct to a json byte slice
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	//print encoded JSON string
	fmt.Println(string(jsonData))

	//creating empty person struct to hold decoded json
	var decodedPerson Person

	//decoding the json data into empty struct
	errr := json.Unmarshal(jsonData,&decodedPerson)
	if errr != nil {
		fmt.Println("Error decoding json:",err)
		return
	}
	fmt.Println(decodedPerson)

}
