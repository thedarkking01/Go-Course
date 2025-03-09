package main

import "fmt"

func main() {
	//  for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// infinite loop
	for {
		fmt.Println("hello")
		break // to prevent infinite loop
	}

	// while loop
	i := 1
    for i <= 5 {
        fmt.Println(i)
        i++
    }

	// with range
	for i:= range 3{
		fmt.Println(i)
	}
}