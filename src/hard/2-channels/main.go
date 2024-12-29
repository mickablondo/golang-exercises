package main

import (
	"fmt"
	"time"
)

// Function to be run as a goroutine
func printNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

/*
-- What do we learn ?
-- How to use channels for communication between goroutines !
-- Channels provide a way for goroutines to synchronize and communicate.
*/
func main() {
	ch := make(chan int)
	go printNumbers(ch)

	for num := range ch {
		fmt.Println(num)
	}

	fmt.Println("Main function finished")
}
