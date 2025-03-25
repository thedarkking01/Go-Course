// A WaitGroup is part of the sync package. It provides a way to wait for a set of goroutines to complete their execution before proceeding.

// A WaitGroup has three main functions:

// Add(int): Adds the number of goroutines to wait for.

// Done(): Decrements the counter by one, signaling that a goroutine is finished.

// Wait(): Blocks until the counter becomes zero, which means all the goroutines have completed.

package main

import (
	"fmt"
	"sync"
	"time"
)
func main() {
	// Create a new WaitGroup
	var wg sync.WaitGroup

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		// Increment the WaitGroup counter by 1
		wg.Add(1)

		go func(i int) {
			defer wg.Done() // Decrement the counter when the goroutine finishes
			fmt.Printf("Goroutine %d is working\n", i)
			time.Sleep(2 * time.Second) // Simulate work
			fmt.Printf("Goroutine %d is done\n", i)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines are done")
}
