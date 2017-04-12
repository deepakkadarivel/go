package main

import "fmt"

// Function accepting 2 int parameters and returning an int
func plus(a int, b int) int {
	return a + b
}

// Function with multiple consecutive parameters of same type
func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 3)
	fmt.Println("Result of plus() with parameter's 1, 3 is : ", res)

	res = plusplus(1, 2, 3)
	fmt.Println("Result of plusplus() with parameter's 1, 2, 3 is : ", res)
}
