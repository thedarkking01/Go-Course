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
        FirstName: "Sourabh",
        LastName:  "Bais",
        Age:       22,
    }

    // Accessing struct fields
    fmt.Println(person1.FirstName) 
    fmt.Println(person1.LastName)  
    fmt.Println(person1.Age)       
}
