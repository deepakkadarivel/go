package main

import "fmt"

func main() {
	// Create an array that will hold exactly 5 ints
	a := [5]int{}
	fmt.Println("a : ", a)

	// Declaring an array variable of type int and length 5
	var b [5]int
	fmt.Println("b : ", b)

	// Setting a value for array element using array[index] = value syntax
	b[2] = 111
	fmt.Println("b values : ", b)

	// using built in len to return the length of an array
	fmt.Println("length of b is : ", len(b))

	// declaring and initializing an array in one line
	c := [5]int{1,2,3,4,5}
	fmt.Println("c : ", c)

	// multi-dimensional array
	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i +  j
		}
	}
	fmt.Println("Multi-dimensional array d : ", d)


}
