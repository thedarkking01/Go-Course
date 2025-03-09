package main

import "fmt"

const DaysInWeek int = 7
const Name string = "Go Programming"

func main() {
    fmt.Println(DaysInWeek)
    fmt.Println(Name)
	const (
		Length = 10
		Width  = 5
		Area   = Length * Width
	)
	fmt.Println(Area)	
}
