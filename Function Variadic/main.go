package main

import "fmt"

// Function accepting variadic number of int arguments
func sums(nums ...int) int {
	fmt.Println("numd : ", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	res := sums(1,2,3)
	fmt.Println("Result for 1, 2, 3 is : ", res)

	// Using slice as an input param to variadic function
	s := []int{4,5,6}
	res = sums(s...)
	fmt.Println("Result for 4, 5, 6 is : ", res)

}
