package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)  // Create a buffered channel with a capacity of 2

	ch <- 1  // Send 1 to the channel (won't block)
	ch <- 2  // Send 2 to the channel (won't block)

	// Main goroutine receives from the channel
	fmt.Println(<-ch)  // 1
	fmt.Println(<-ch)  // 2
}
