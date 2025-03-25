//You can define methods (functions tied to a struct).

package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Car struct {
	Brand string
	Model string
}

//sayHello() is a method that belongs to Student.
func (s Student) sayHello() {
	fmt.Println("Hello, my name is", s.Name)
}

//Instead of methods, you can use regular functions that accept a struct
func PrintCarInfo(c Car) {
	fmt.Printf("Car: %s %s\n", c.Brand, c.Model)
}

func main() {
	student1 := Student{"Shahid", 18}
	student1.sayHello()

	car := Car{"Porsche", "911 GT3 RS"}
	PrintCarInfo(car)

	/*
		✅ When to Use Methods vs Functions?

		Use methods if the function belongs to the struct (e.g., Greet() for Student).

		Use functions if the struct is just data passed around (e.g., PrintCarInfo() for Car).
	*/
}
