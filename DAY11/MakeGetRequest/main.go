package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	performGetRequest()
}

func performGetRequest() {

	const url = "https://www.wikipedia.org/"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status :", response.Status)
	fmt.Println("Status code :", response.StatusCode)
	fmt.Println("content length :", response.ContentLength)
	// content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(content)
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	bytecount,_ := responseString.Write(content)

	fmt.Println(bytecount)
}
