package main

import "fmt"

//shorthand syntax wont work in global
const age = 30

func main() {
	const name string = "golang"

	// name = "javascript" // This will cause an error because constants cannot be reassigned

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// Constants can be declared in a group
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
