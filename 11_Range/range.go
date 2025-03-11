package main

import "fmt"

func main() {
    // Range with a slice
    slice := []int{1, 2, 3, 4}
    for i, v := range slice {
        fmt.Printf("Index: %d, Value: %d\n", i, v)
    }

    // Range with a map
    m := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range m {
        fmt.Printf("Key: %s, Value: %s\n", k, v)
    }

    // Range with a string
    s := "hello"
    for i, r := range s {
        fmt.Printf("Index: %d, Rune: %c\n", i, r)
    }

    
}
