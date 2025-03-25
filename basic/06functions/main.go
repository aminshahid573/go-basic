package main

import "fmt"

func greet() {
	fmt.Println("Hello, Welcome to GO functions")
}

//function with parameters
func greetWithName(name string) {
	fmt.Println("Hello,", name)
}

//function with return value
func add(a int, b int) int {
	return a + b
}

//TODO: Assignments
func square(a int) int {
	return a * a
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	//function without parameters
	greet()
	//function with parameters
	greetWithName("Shahid Amin")
	//function with return
	fmt.Println("Sum : ", add(5, 9))

	//Assignments
	fmt.Println("Square : ", square(5))
	fmt.Print("Is Even : ", isEven(5))
}
