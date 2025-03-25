package main

import "fmt"

func main() {
	var num int = 50
	ptr := &num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at pointer:", *ptr)

}
