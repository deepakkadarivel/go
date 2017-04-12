package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroPointer(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial : ", i)

	// Here i is just a parameter the value will not change
	zeroval(i)
	fmt.Println("zeroval : ", i)

	// Here &i is memory address of i and value inside address will change
	zeroPointer(&i)
	fmt.Println("zero-pointer : ", i)
}
