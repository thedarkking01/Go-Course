// package main

// import "fmt"

// // Function with parameters
// func add(a int, b int) int {
//     return a + b
// }

// func main() {
//     sum := add(3, 5) // Calling the function with arguments
//     fmt.Println("Sum:", sum) // Prints: Sum: 8
// }

package main

import "fmt"

// Function with multiple return values
func divide(a int, b int) (int, string) {
    if b == 0 {
        return 0, "Error: Division by zero"
    }
    return a / b, "Success"
}

func main() {
    result, message := divide(10, 2)
    fmt.Println("Result:", result, "Message:", message) // Prints: Result: 5 Message: Success

    result, message = divide(10, 0)
    fmt.Println("Result:", result, "Message:", message) // Prints: Result: 0 Message: Error: Division by zero
}
