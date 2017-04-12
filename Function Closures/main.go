package main

import "fmt"

// Function returns a function which is defined as a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// Assigning a closure function to a variable which will act as a function
	nextInt := intSeq()

	fmt.Println("Iteration 1 : ", nextInt())
	fmt.Println("Iteration 2 : ", nextInt())
	fmt.Println("Iteration 3 : ", nextInt())

	// Re-initializing and checking the value of i
	nextLevelInt := intSeq()
	fmt.Println("Iteration n : ", nextLevelInt())
}

