package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	res := fact(7)
	fmt.Println("Factorial of 7 is : ", res)
}
