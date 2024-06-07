package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is file handling....!!!!")
	content := "This is content that needs to go in a file"

	file, err := os.Create("./myFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilError(err)


	fmt.Println("length is ", length)
	readFile("./myFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("content in the file is \n", string(databyte))

}

func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}
