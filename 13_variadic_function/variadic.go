package main

import "fmt"

func sum(nums ...int) int {
	//we got nums as slice of int
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

//if we want any type then we use ...interface like
// func sum(nums ...interface{}) int {}

func main() {
	//accepts n no of parameter
	// fmt.Println(1, 2, 3, "Hello")

	// result := sum(3, 4, 6, 8)
	//we can also use slice
	nums := []int{3, 4, 6}
	result := sum(nums...)
	fmt.Println(result)

}
