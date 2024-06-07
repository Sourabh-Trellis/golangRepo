package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple","mango","banana"}

	fmt.Printf("type of slice is %T.\n", fruits)
	fmt.Println(fruits)
	fmt.Println("cap:",cap(fruits))

	fruits = append(fruits, "peach","kiwi")
	fmt.Println(fruits)
	fmt.Println("cap:",cap(fruits))


	fruits = append(fruits[:4])
	fmt.Println(fruits)
	fmt.Println("cap:",cap(fruits))

	highscores := make([]int, 5)
	highscores[0]=119
	highscores[1]=115
	highscores[2]=118
	highscores[3]=112
	highscores[4]=111
	
	highscores = append(highscores, 199, 299)

	fmt.Println(highscores)

	sort.Ints(highscores)

	fmt.Println(highscores)

	uphighscores := append(highscores[:3], highscores[4:]...)
	fmt.Println(uphighscores)

	

}
