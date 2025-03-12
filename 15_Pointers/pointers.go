package main
import "fmt"

func increment(x *int) {
    *x++  // Dereference the pointer to modify the value
}

func main() {
    a := 10
    increment(&a)  // Pass the address of 'a' to the function
    fmt.Println(a)  // Output: 11

	x := 10
    ptr := &x  // pointer to x

    fmt.Println("Address of x:", ptr)  // prints address of x
    fmt.Println("Value of x:", *ptr)   // prints value of x (dereferencing pointer)

    *ptr = 20  // change the value of x through the pointer
    fmt.Println("New value of x:", x)
}
