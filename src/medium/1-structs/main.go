package main

import "fmt"

// Define a struct type
type Person struct {
	Name string
	Age  int
}

/*
 -- What do we learn ?
 -- How to define and use structs !
 -- Structs are a way to group related data together.
*/
func main() {
	// Create an instance of the struct
	person := Person{Name: "MiK@", Age: 40}

	// Access and print struct fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}
