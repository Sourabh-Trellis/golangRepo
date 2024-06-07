package main

import (
	"fmt"
	"regexp"
)

func main() {

	//creating regexp object can be reused later
	re, _ := regexp.Compile("geek")

	//string to be matched
	str := "I love geeksforgeeks"

	//finding index of regexp in string
	match := re.FindStringIndex(str)
	fmt.Println(match)

	str2 := "I am computer Engineer"
	match2 := re.FindStringIndex(str2)
	fmt.Println(match2)

	//creating a regexp pattern
	reg2, _ := regexp.Compile("[0-9]+-v.*g")
	//matching one or more numbers followed by v and any number of characters upto 'g'
	match3 := reg2.FindString("200034-vani-gupta")
	fmt.Println(match3)

	match4 := re.FindAllStringSubmatchIndex("I am a geek at geekforgeeks", -1)
	fmt.Println(match4)

	reg3,_ := regexp.Compile(" ")
	match5 := reg3.ReplaceAllString("All i do is code all time.","+")
	fmt.Println(match5)

}
