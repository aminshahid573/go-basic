package main

import "fmt"

//if parameters are of same type we can just use comma like
// func add(a, b int) int {}
func add(a int, b int) int {
	return a + b
}

//go function can return multiple value
func getLanguages() (string, string, string) {
	return "Go", "Java", "C++"
}

//function in go is first class citizen
//function can be assigned to a variable

//passing function as parameter
func processIt(fn func(a int) int) {
	val := fn(2)
	fmt.Println(val)
}

//return function
func returnFunc() func(a int) int {
	return func(a int) int {
		return a + 1
	}
}

func main() {
	result := add(3, 5)
	fmt.Println(result)

	//calling getLanguages
	l1, l2, l3 := getLanguages()
	fmt.Println(l1, l2, l3)

	//calling processIt
	fn := func(a int) int {
		return a + 1
	}
	processIt(fn)

	//calling re
	fn2 := returnFunc()
	fmt.Println(fn2(2))
}
