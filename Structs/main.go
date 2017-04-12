package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	// Creating a new struct value
	fmt.Println("person 1 : ", person{"Deepak", 23})

	// Naming the fields during initialization of struct
	fmt.Println("person 2 : ", person{name: "Abhi", age: 19})

	// Omitting a filed that will be zero-valued
	fmt.Println("person 3 : ", person{name: "Sagar"})

	// Accessing struct fields with dot
	s := person{"Rishi", 30}
	fmt.Printf("name : %s, age : %d \n", s.name, s.age)

	// Creating a variable that references to struct via a pointer
	sp := &s
	fmt.Println("Pointed reference of s is : ", sp.name)

}
