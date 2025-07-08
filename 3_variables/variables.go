package main

import "fmt"

func main() {
	// String
	var name string = "Shahid"
	var name2 = "Amin" // this is also valid
	fmt.Println(name, name2)

	//boolean
	var isAdult bool = true
	fmt.Println(isAdult)

	//integer
	var age int = 30
	fmt.Println(age)

	//float
	var height float64 = 5.9
	fmt.Println(height)

	//variable with no value
	var country string
	fmt.Println(country)

	//Shorthand declaration
	//best for declaring variable and assign value immediately
	city := "Biratnagar"
	fmt.Println(city)

}
