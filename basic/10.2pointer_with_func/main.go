/*
By default, Go passes arguments by value, meaning a function gets a copy of the variable.
Using pointers allows modifying the original value.
*/

package main

import "fmt"

//without pointers (pass by value)
func updateValue(x int) {
	x = 100
}

// with pointers (pass by refrence)
func updateValueUsingPtr(x *int) {
	*x = 100 // Modify the value at the pointer address
}
func main() {
	//without pointer
	num := 50
	updateValue(num)
	fmt.Println("Value after function call:", num) //❌ Still 50 (no change)

	updateValueUsingPtr(&num)
	fmt.Println("Value after function call:", num) // ✅ Now 100

}
