package main

import (
	"fmt"
	"reflect"
)

func printTypeAndValue(v interface{}) {

	//get type of the variable
	t := reflect.TypeOf(v)
	fmt.Println("Type:", t)

	//get value of the variable
	value := reflect.ValueOf(v)

	// fmt.Println(reflect.Slice,reflect.Map)

	//check if variable is a map or slice
	if t.Kind() == reflect.Slice {

		//iterate over the elements of slice and print each element
		fmt.Println("value:")
		for i := 0; i < value.Len(); i++ {
			fmt.Println("\t", value.Index(i))
		}
	} else if t.Kind() == reflect.Map {
		//iterate over the elements of the map and print each key-value
		fmt.Println("Value(Map):")
		keys := value.MapKeys()
		for _, key := range keys {
			val := value.MapIndex(key)
			fmt.Println("\t", key, ":", val)
		}
	} else {
		//print the value of the variable
		fmt.Println("value:", value)
	}
}

func main() {
	i := 2
	printTypeAndValue(i)

	v := []int{20, 58, 35, 59, 02, 07}
	printTypeAndValue(v)

	m := map[string]int{"a":1,"b":2,"c":3}
	printTypeAndValue(m)
}
