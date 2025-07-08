package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5

	//break keyword is not necessary
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: //optional
		fmt.Println("other")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	//type switch
	// interface{} means any type
	whoAmI := func(i interface{}) {
		switch t := i.(type) { //we can also write `switch i.(type)` if we dont need to store the type
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Printf("other type : %T", t)
		}
	}

	whoAmI(42.21)

}
