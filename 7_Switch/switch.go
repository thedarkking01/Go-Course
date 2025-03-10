package main

import "fmt"

func main() {
	// normal switch

	// day:=3
	// switch day {
	// 	case 1:
	//         println("Monday")
	//     case 2:
	//         println("Tuesday")
	//     case 3:
	//         println("Wednesday")
	//     case 4:
	//         println("Thursday")
	//     case 5:
	//         println("Friday")
	//     case 6:
	//         println("Saturday")
	//     case 7:
	//         println("Sunday")
	//     default:
	//         println("Invalid day of the week")
	// }

	//switch with no condition
	x := 10

	switch {
	case x > 5:
		fmt.Println("x is greater than 5")
	case x == 5:
		fmt.Println("x is equal to 5")
	default:
		fmt.Println("x is less than 5")
	}

	//multiple condition
	day := "Tuesday"

    switch day {
    case "Monday", "Tuesday", "Wednesday":
        fmt.Println("Start of the week")
    case "Thursday", "Friday":
        fmt.Println("End of the week")
    default:
        fmt.Println("Weekend!")
    }

	//type switch
	var y interface{} = 42

    switch v := y.(type) {
    case int:
        fmt.Println("x is an int with value", v)
    case string:
        fmt.Println("x is a string with value", v)
    default:
        fmt.Println("Unknown type")
    }
}