package main

import "fmt"

func main() {
	//array is fixed size collection of elements of the same type
	// memory optimization
	// very fast to access
	var nums [4]int

	fmt.Println(len(nums)) // Length of the array

	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	var name [3]string
	name[0] = "John"
	name[1] = "Doe"
	name[2] = "Smith"

	fmt.Println(name)

	//another way to declare and initialize an array
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// Multidimensional array
	multiDimension := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(multiDimension)
}
