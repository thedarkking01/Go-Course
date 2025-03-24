package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Second) // Sleep for 1 second
	}
}

func main() {
	// Start a goroutine
	go printMessage("Hello from goroutine")

	// This will be executed in the main goroutine
	printMessage("Hello from main")

	// Sleep to allow goroutines to finish
	time.Sleep(6 * time.Second)
}
