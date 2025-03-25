package main

import "fmt"

func main() {
	var numbers [5]int //Array of 5 integers
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println(numbers)

	numbersSlices := []int{10, 20, 30}

	fmt.Println("Original Slice: ", numbersSlices)

	numbersSlices = append(numbersSlices, 40)
	fmt.Println("Modified Slice: ", numbersSlices)

	//loping through arrays and slices
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

}
