package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("web request")

	resp,err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("type of responce is %T\n",resp)
	defer resp.Body.Close()

	databytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
 