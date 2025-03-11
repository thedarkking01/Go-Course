package main

import "fmt"

func main() {
    // Creating a map using make
    m := make(map[string]int)

    // Adding elements to the map
    m["apple"] = 5
    m["banana"] = 3
    m["cherry"] = 7

    // Updating a value
    m["apple"] = 10

    // Accessing a value
    fmt.Println("Value for apple:", m["apple"]) // Prints: 10

    // Checking if a key exists
    value, exists := m["banana"]
    if exists {
        fmt.Println("Value for banana:", value) // Prints: 3
    } else {
        fmt.Println("Banana not found")
    }

    // Deleting an element
    delete(m, "cherry")
    fmt.Println("After deleting cherry:", m) // Prints: map[apple:10 banana:3]

    // Iterating over the map
    fmt.Println("Iterating over the map:")
    for key, value := range m {
        fmt.Println(key, value)
    }

    // Getting the length of the map
    fmt.Println("Length of the map:", len(m)) // Prints: 2
}
