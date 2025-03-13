package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Initialize a struct instance
    person1 := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
    }

    // Accessing struct fields
    fmt.Println(person1.FirstName) // Output: John
    fmt.Println(person1.LastName)  // Output: Doe
    fmt.Println(person1.Age)       // Output: 30
}
