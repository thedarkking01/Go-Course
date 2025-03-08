package main

import "fmt"

func main() {
    // Variable declarations with type
    var age int = 25
    var name string = "Alice"

    // Short variable declaration
    height := 5.6
    isActive := true

    // Multiple variable declaration
    var x, y int = 10, 20
    var firstName, lastName = "John", "Doe"

    // Zero value declaration
    var city string // zero value is ""

    fmt.Println(age, name, height, isActive)
    fmt.Println(x, y, firstName, lastName, city)
}
