package main

import "fmt"

func main() {
    x := 10
    if x > 15 {
        fmt.Println("x is greater than 15")
    } else if x > 5 {
        fmt.Println("x is greater than 5 but less than or equal to 15")
    } else {
        fmt.Println("x is less than or equal to 5")
    }

	// Shortened if statement -> you can also declare and initialize a variable within the if statement.
	if x := 10; x > 5 {
        fmt.Println("x is greater than 5")
    }
}
