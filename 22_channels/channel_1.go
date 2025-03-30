// package main

// import "fmt"

// func main() {
//     // Create a new channel of integers
//     ch := make(chan int)

//     // Start a new goroutine
//     go func() {
//         // Send a value into the channel
//         ch <- 42
//         fmt.Println("Value sent to channel.")
//     }()

//     // Receive the value from the channel
//     receivedValue := <-ch
//     fmt.Println("Received from channel:", receivedValue)
// }
