package main

import "fmt"

func main() {
    // Creating a slice using a literal
    numbers := []int{10, 20, 30, 40, 50}
    
    // Accessing elements
    fmt.Println("Second element:", numbers[1])  // Prints: 20
    
    // Slicing the slice
    sub := numbers[1:4]  // Get a sub-slice from index 1 to 3
    fmt.Println("Sub-slice:", sub)  // Prints: [20 30 40]
    
    // Appending to the slice
    numbers = append(numbers, 60, 70)
    fmt.Println("After appending:", numbers)  // Prints: [10 20 30 40 50 60 70]
    
    // Getting length and capacity
    fmt.Println("Length:", len(numbers))  // Prints: 7
    fmt.Println("Capacity:", cap(numbers))  // Prints the capacity (can vary)

	// Creating a 2D slice (3x3 matrix)
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    // Accessing elements
    fmt.Println("Element at row 1, column 2:", matrix[0][1])  // Prints: 2

    // Modifying an element
    matrix[2][2] = 10
    fmt.Println("Modified matrix:", matrix)  // Prints: [[1 2 3] [4 5 6] [7 8 10]]

    // Iterating over 2D slice
    fmt.Println("Iterating over matrix:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
        }
    }
}
