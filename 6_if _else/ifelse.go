package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else {
	// 	fmt.Println("You are not an adult.")
	// }

	percentage := 40

	if percentage > 65 {
		fmt.Println("B+")
	} else if percentage > 80 {
		fmt.Println("A")
	} else if percentage <= 65 {
		fmt.Println("Fail")
	} else {
		fmt.Println("A+")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("User has admin permission")
	}

	//we can declare variable in if statement
	if age := 15; age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	//go doesnot have ternary operator, you will have to use normal if else

}
