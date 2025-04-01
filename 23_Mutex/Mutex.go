package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex   // Mutex for synchronization
var counter int = 0  // Shared resource

// Function that modifies the shared resource
func increment(wg *sync.WaitGroup) {
	defer wg.Done()  // Notify when the goroutine is done
	mu.Lock()         // Acquire the lock
	counter++         // Critical section: accessing shared resource
	mu.Unlock()       // Release the lock
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for goroutines to finish

	// Start 10 goroutines to increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)            // Increment the WaitGroup counter
		go increment(&wg)    // Call increment function in a new goroutine
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter:", counter) // Print final value of counter
}
