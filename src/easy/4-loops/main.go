package main

import "fmt"

/*
 -- What do we learn ?
 -- How to loop !
 -- Hey mum, look at this carousel !
*/
func main() {
	for i := 0; i < 3; i++ {
		if i < 2 {
			fmt.Println("Wooo ... looping ", i+1)
		} else {
			fmt.Println("Stooop ... i want to vomit ...")
		}
	}

	fmt.Println("\n How do you write your name ?")
	name := "MiK@"
	for i, c := range name {
		fmt.Println("Letter ", i+1, " is ", string(c))
	}
}
