package main

import "fmt"

/*
 -- What do we learn ?
 -- How to use simples operators !
 -- Hmmm ... okay ...
*/
func main() {
	x := 5
	y := 4

	fmt.Println("Addition : ", x, "+", y, "=", x+y)
	fmt.Println("Substraction : ", x, "-", y, "=", x-y)
	fmt.Println("Multiplication : ", x, "x", y, "=", x*y)
	fmt.Println("Division : ", x, "/", y, "=", float64(x)/float64(y))
	fmt.Println("Modulo : ", x, "%", y, "=", x%y)
}
