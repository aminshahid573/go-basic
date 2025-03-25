package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func changeName(p *Person, newName string) {
	p.Name = newName
}

func main() {
	person := Person{"Shahid", 18}
	changeName(&person, "Zeeshan")
	fmt.Println("Updated Name:", person.Name) // ✅ Name changed to "Zeeshan"

	/*
	✅ Why use a pointer to a struct?

		Without a pointer, the function would receive a copy of Person (no changes).

		With a pointer, it modifies the original struct.
	*/
}
