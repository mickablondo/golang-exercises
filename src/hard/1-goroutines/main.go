package main

import (
	"fmt"
	"time"
)

// Function to be run as a goroutine
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

/*
-- What do we learn ?
-- How to use goroutines !
-- Goroutines are lightweight threads of execution.
*/
func main() {
	// Start a new goroutine
	go printNumbers()

	// Main goroutine continues to run
	fmt.Println("Goroutine started")

	// Wait for the goroutine to finish
	time.Sleep(6 * time.Second)
	fmt.Println("Main function finished")
}
