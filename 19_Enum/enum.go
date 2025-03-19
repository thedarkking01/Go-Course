package main

import "fmt"

// Define the enum using iota
type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    fmt.Println(Sunday)   // Output: 0
    fmt.Println(Monday)   // Output: 1
    fmt.Println(Tuesday)  // Output: 2
    fmt.Println(Wednesday) // Output: 3

    // You can also create a function to display the enum as a string
    fmt.Println(dayString(Sunday)) // Output: Sunday
}

func dayString(d Day) string {
    switch d {
    case Sunday:
        return "Sunday"
    case Monday:
        return "Monday"
    case Tuesday:
        return "Tuesday"
    case Wednesday:
        return "Wednesday"
    case Thursday:
        return "Thursday"
    case Friday:
        return "Friday"
    case Saturday:
        return "Saturday"
    default:
        return "Unknown"
    }
}
