/*
A struct in Go is a collection of fields that group related data together. Think of it like an object in JavaScript or a class without methods in Python.
*/

/*
type StructName struct {
    Field1 DataType
    Field2 DataType
}
*/

package main

import "fmt"

//Define a struct
type Student struct {
	Name  string
	Age   int
	Grade string
}

func main() {
	//creating a struct instance
	//Method 1: initializing using values directly
	//field order matter
	student1 := Student{"Shahid", 18, "A"}

	// Method 2: initializing Using Named Fields (Recommended)
	//field order doesn't matter
	student2 := Student{
		Name:  "Zeeshan",
		Age:   18,
		Grade: "A+",
	}

	//Method 3: initializing Using var with Default Zero Values
	var student3 Student
	student3.Name = "Rajan"
	student3.Age = 19
	student3.Grade = "B+"

	//Accessing struct fields
	fmt.Println("Name: ", student1.Name)
	fmt.Println("Age: ", student1.Age)
	fmt.Println("Grade: ", student1.Grade)

	fmt.Println("student1: ", student1)
	fmt.Println("student2: ", student2)
	fmt.Println("student3: ", student3)
}
