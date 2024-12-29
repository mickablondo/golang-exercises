package main

import "fmt"

// Define a struct type
type Rectangle struct {
	Width, Height int
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

/*
 -- What do we learn ?
 -- How to define and use methods on structs !
 -- Methods allow us to associate functions with types.
*/
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of rectangle:", rect.Area())
}
