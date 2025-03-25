package main

import "fmt"

func main() {

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	fmt.Printf("Multiplication table for %d:\n", num)

	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}
