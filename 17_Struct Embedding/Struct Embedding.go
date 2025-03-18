package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Define another struct 'Employee' that embeds 'Person'
type Employee struct {
    Person  // This embeds the 'Person' struct
    Position string
}

func main() {
    // Initialize an 'Employee' instance
    employee1 := Employee{
        Person: Person{
            FirstName: "Sourabh",
            LastName:  "Bais",
            Age:       22,
        },
        Position: "Software Engineer",
    }

    // Accessing embedded struct fields
    fmt.Println(employee1.FirstName)  
    fmt.Println(employee1.LastName)  
    fmt.Println(employee1.Age)        
    fmt.Println(employee1.Position) 
}
