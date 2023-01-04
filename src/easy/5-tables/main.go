package main

import "fmt"

/*
 -- What do we learn ?
 -- How to make matrices !
 -- So cool, i love theses movies !
*/
func main() {
	// init array of 10 int
	var my10IntArray [10]int

	// loop array
	for i := 0; i < len(my10IntArray); i++ {
		my10IntArray[i] = i
	}

	for index, value := range my10IntArray {
		fmt.Println("my10IntArray[", index, "] = ", value)
	}

	// init array with multiple int
	myIndefineIntArray := [...]int{1, 2, 3}
	fmt.Println("myIndefineIntArray = ", myIndefineIntArray)
}
