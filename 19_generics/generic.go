package main

import "fmt"

//you can also do [T any]
// or
// func printSlice[T int | string | bool](items []T) {}

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type stack[T any] struct {
	elements []T
}

func main() {

	// myStack := stack[string]{
	// 	elements: []string{"golang", "java"},
	// }
	// fmt.Println(myStack)

	// printSlice(nums)

	// nums := []int{1, 2, 3}
	// names := []string{"golang", "typescript"}
	booleans := []bool{true, false, false}
	printSlice(booleans)
}
