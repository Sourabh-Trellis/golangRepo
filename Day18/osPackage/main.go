package main

import (
	"fmt"
	"os"
)

func main() {
	//os package

	//get current working directory
	curdir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println("Current directory :", curdir)

	//creating new directory
	dirName := "new_directory"

	//check if directory already exists
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		//directory does not exist,create it
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Println("Error :",)
			return
		}
		fmt.Println("Directory created successfully.")
	} else if err != nil {
		//if some other error occurred
		fmt.Println("Error:",err)
		return
	} else {
		//Directory already exists
		fmt.Println("Directory already exists")
	}
}
