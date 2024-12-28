package main

import "fmt"

// Function that adds two integers and returns the result
func add(a int, b int) int {
	return a + b
}

// Function that multiplies two integers and returns the result
func multiply(a int, b int) int {
	return a * b
}

/*
 -- What do we learn ?
 -- How to use functions !
 -- So cool, i can make my code reusable!
*/
func main() {
	sum := add(3, 4)
	product := multiply(3, 4)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
}
