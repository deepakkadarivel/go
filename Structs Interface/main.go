package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	peri() float64
}

type rect struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) peri() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi + (c.radius  * c.radius)
}

func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g : ", g )
	fmt.Println("Area : ", g.area())
	fmt.Println("Peri : ", g.peri())
}

func main() {
	r := rect{3,4}
	c := circle{5}

	measure(r)
	measure(c)
}
