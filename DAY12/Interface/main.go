package main

import "fmt"

// creating an interface
type tank interface {

	//methods
	total_area() float64
	volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

type myval2 struct {
	radius float64
	height float64
}

func (m myvalue) total_area() float64 {
	return 2*m.radius*m.radius +
		2*3.14*m.radius*m.radius
}

func (m myvalue) volume() float64 {
	return 3.14 * m.radius * m.radius
}

func (m myval2) total_area() float64 {
	return 2*m.radius*m.radius +
		2*3.14*m.radius*m.radius
}

func (m myval2) volume() float64 {
	return 3.14 * m.radius * m.radius
}

func main() {

	var t tank = myvalue{10, 40}
	fmt.Println(t.total_area())

	var T = myval2{10, 10}

}
