package main
import "fmt"

func main() {
    // Declaring and initializing an array
    var arr [5]int = [5]int{1, 2, 3, 4, 5}

    // Accessing array elements
    fmt.Println("First element:", arr[0])
    fmt.Println("Third element:", arr[2])

    // Modifying an array element
    arr[4] = 10
    fmt.Println("Modified array:", arr)

    // Iterating over an array using for loop
    fmt.Println("Array elements using for loop:")
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // Iterating using range
    fmt.Println("Array elements using range:")
    for index, value := range arr {
        fmt.Printf("Index %d: Value %d\n", index, value)
    }
}
