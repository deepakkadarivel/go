package main

import "fmt"

type rect struct {
	width int
	height int
}

// Method has a receiver type of *rect
func (r *rect) area() int  {
	return r.width * r.height
}

// Method has a value receiver type rect
func (r rect) peri() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 6}

	fmt.Println("Area : ", r.area())
	fmt.Println("Peri : ", r.peri())

	rec := &r
	fmt.Println("Area 2 : ", rec.area())
	fmt.Println("Peri 2 : ", rec.peri())
}
