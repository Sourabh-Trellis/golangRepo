package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"Tag,omitempty"`
}

func main() {
	fmt.Println("welcome")
	//encodeJSON()
	decodeJSON()
}

func encodeJSON() {

	courses := []course{
		{"ReactJS Bootcamp", 299, "linkedin", "asdfg12345", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 2111, "udemy", "asd345", []string{"full-stack", "js", "reactjs"}},
		{"ReactJS Bootcamp", 299, "coursera", "sdsdsd", nil},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJSON() {

	jsondatafromweb := []byte(`
		{
			"Coursename": "MERN Bootcamp",
			"Price": 2111,
			"website": "udemy",
			"Tag": ["full-stack","js","reactjs"]
	} 
	`)

	var Course course
	checkValid := json.Valid(jsondatafromweb)

	if checkValid {
		fmt.Println("JSON")
		json.Unmarshal(jsondatafromweb, &Course)
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("not valid json data")
	}
}
