package main

import "fmt"

//for -> only contruct in go looping
func main() {
	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	//classic for loop
	// also can use continue and break
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// range loop
	for i := range 3 {
		fmt.Println(i)
	}
}
