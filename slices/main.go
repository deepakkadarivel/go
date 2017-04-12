package main

import "fmt"

func main() {

	// Using builtin make create an empty slice with non-zero length
	s := make([]string, 3)
	fmt.Println("Slice s : ", s)
	fmt.Println("Length of slice s : ", len(s))
	fmt.Println("Capacity of slice s is : ", cap(s))

	// Assigning values to slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Assiged values to s is : ", s)
	fmt.Println("Length of slice s : ", len(s))
	fmt.Println("Capacity of slice s is : ", cap(s))

	// Appending values to slice
	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f", "g", "h")
	fmt.Println("Appended values to slice is : ", s)
	fmt.Println("Length of slice is : ", len(s))
	fmt.Println("Capacity of slice s is : ", cap(s))

	// create a new slice
	c := make([]string, len(s))

	// Copy a slice into another slice
	copy(c, s)
	fmt.Println("Slice s : ", s)
	fmt.Println("Slice c ", c)

	// Slicing a slice
	var a = s[:2]
	fmt.Println("s[:2] is : ", a)

	var b = s[2:]
	fmt.Println("s[2:] is : ", b)

	// Declaring a slice and initializing the value
	sl := []string{"x","y","z"}
	fmt.Println("Slice sl is : ", sl)
}
