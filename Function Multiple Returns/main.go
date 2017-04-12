package main

import "fmt"

// Function signature shows it return 2 ints, (int, int)
func vals() (int, int) {
	return 2, 3
}

func main() {
	// Here the return values are assigned  to 2 different values.
	a, b := vals()
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	// subset of return values
	_, c := vals()
	fmt.Println("c : ", c)


}
