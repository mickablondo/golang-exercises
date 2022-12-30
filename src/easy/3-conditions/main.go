package main

import (
	"fmt"
	"runtime"
)

/*
-- What do we learn ?
-- Using if, else, switch case and import multiples packages !
-- Yeah !
*/
func main() {
	bool1 := true
	bool2 := false

	if bool1 {
		fmt.Println("bool1 is true !")
	}

	if !bool2 {
		fmt.Println("bool2 is false !")
	}

	if 1 > 2 {
		fmt.Println("you're in an alternative universe ...")
	} else if 1 != 1 {
		fmt.Println("where are you ?")
	} else {
		fmt.Println("Else is OK !")
	}

	os := runtime.GOOS // inspired by https://go.dev/tour/flowcontrol/9
	fmt.Print("My OS is ... ")
	switch os {
	case "windows":
		fmt.Println("Windows.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("neither windows nor linux but ", os)
	}
}
