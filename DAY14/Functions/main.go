package main

import "fmt"

type rectangle struct {
	length  float32
	breadth float32
}

// NORMAL FUNCTION
// func area(length, breadth float32) float64 {
// 	return float64(length) * float64(breadth)
// }

// RECIEVER FUNCTION
func (r rectangle) area() float32 {
	return r.length * r.breadth
}

// POINTER RECIEVER FUNCTION
func (r *rectangle) setLength(l float32) {
	r.length = l
}

// POINTER RECIEVER FUNCTION
func (r *rectangle) setBreadth(b float32) {
	r.breadth = b
}

func main() {

	rect := rectangle{}
	rect.setLength(4)
	rect.setBreadth(5)

	// area := area(rect.breadth, rect.length)

	area := rect.area()
	fmt.Println(area)
}
