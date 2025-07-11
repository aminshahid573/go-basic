package main

import "fmt"

//call by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum:", num)
// }

//call by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum:", *num)
}

func main() {
	//call by value part
	// num := 1
	// changeNum(num)
	// fmt.Println("After changeNum in main:", num)

	//by reference part
	num := 1
	fmt.Println("Memory address", &num)

	changeNum(&num)
	fmt.Println("After changeNum in main:", num)
}
