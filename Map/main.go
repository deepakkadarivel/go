package main

import "fmt"

func main() {
	// Create a map with builtin make
	m := make(map[string]int)

	// assigning values to map using m[key] = value synatx
	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map m is : ", m)
	fmt.Println("length of map is : ", len(m))

	// accessing value using a key
	v1 := m["k1"]
	fmt.Println("Value of key 1 is : ", v1)

	// deleting a value by key using builtin delete
	delete(m, "k2")
	fmt.Println("map m after deleting : ", m)

	// checking if a map contains value or not
	v, prs := m["k2"]
	fmt.Println("Value of k2 is : ", v, prs)

	// declaring and initializing a map in a single line
	n := map[int]string{1:"m1", 2:"m2"}
	fmt.Println("new map is : ", n)
}
