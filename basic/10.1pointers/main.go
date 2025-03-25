/*
var ptr *int  // Declaring a pointer to an int
ptr = &num    // Assigning the address of num to ptr

*int → Declares a pointer to an integer.
&num → Gets the memory address of num.
*ptr → Dereferences the pointer (gets the value at that address).
*/

package main

import "fmt"

func main() {
	num := 42
	ptr := &num //pointer stores the address of num

	fmt.Println("Value of num:", num)
	fmt.Println("Memory address of num:", ptr)
	fmt.Println("Value at pointer:", *ptr) //Dereferencing

	// Modifying value using pointer
	*ptr = 10
	fmt.Println("Updated num:", num)

	//A nil pointer is a pointer that doesn’t point to any memory address.
	var nilPtr *int
	if nilPtr == nil {
		fmt.Println("Pointer is nil") // ✅ It prints this
	}

	//🚨 Avoid dereferencing a nil pointer (*ptr)—it will cause a runtime error.

}
