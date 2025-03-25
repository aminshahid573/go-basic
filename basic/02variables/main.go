package main

import "fmt"

func main() {

	//variables
	var name string
	var age int
	var height float64
	// var isStudent bool

	//constant
	const pi = 3.14

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter your height(meter): ")
	fmt.Scanln(&height)

	ageInMonths := age * 12

	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Your are %d yesr old, which is equivalent to %d months.\n", age, ageInMonths)

	fmt.Printf("Your height is %.2f meters.\n", height)

}
