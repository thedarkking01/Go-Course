package main

import "fmt"
func sum(nums ...int) int {
	var sum int = 0
    for _, value := range nums{
        sum += value
    }
    return sum
}

func main() {
	nums :=[]int{3,4,5,6,7}
	result := sum(nums...)
	fmt.Println("Sum:", result)
}