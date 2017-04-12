package main

import "fmt"

func main() {
	// Using range to sum numbers in a slice
	num := []int{1,2,3,4,5}
	sum := 0
	fmt.Println("Sum value is : ", sum)

	for _, val := range num {
		sum += val
	}
	fmt.Println("Sum value after range is : ", sum)

	// Using range to print index of slice
	for i, v := range num {
		fmt.Printf("index : %d = val : %d \n",i, v )
	}

	// Using range to print key value from maps
	mp := map[int]int{0:1, 1:2, 2:3, 3:4}
	for k, v := range mp {
		fmt.Printf("key : %d = value : %d \n", k, v)
	}

	// Using range to access only key from map
	for k, _ := range mp {
		fmt.Printf("Key is : %d \n", k)
	}

	// Range on string iterates over unicode point
	for i, c := range "test" {
		fmt.Printf("index : %d, byte code : %v, string : %s \n", i, c, string(c))
	}


}
