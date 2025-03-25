// a map stores key-value pairs where each key is unique

/*
map[KeyType]ValueType {
	key1: value1,
	key2: value2
 	...
 }
*/

package main

import "fmt"

func main() {
	student := map[string]int{
		"Shahid": 20,
		"Amin":   21,
		"Rahul":  22,
		"Riya":   23,
	}

	fmt.Println("Students Scores: ", student)

	//If you don’t want to initialize values immediately, use make().

	// mapVar := make(map[keyType]valueType)

	studentsMapUsingMake := make(map[string]int)

	studentsMapUsingMake["Zeeshan"] = 20
	studentsMapUsingMake["Aayan"] = 30

	fmt.Println("Students Scores: ", studentsMapUsingMake)

	//Genting Value from map

	zeeshanMarks := studentsMapUsingMake["Zeeshan"]

	fmt.Println("Zeeshan Marks: ", zeeshanMarks)
	//If the key doesn’t exist, it returns the zero value of the data type (e.g., 0 for int, "" for string).

	//Updating Zeeshan Marks
	studentsMapUsingMake["Zeeshan"] = 50
	fmt.Println("Updated Zeeshan Marks: ", studentsMapUsingMake["Zeeshan"])

	//Adding a new key-value pair
	studentsMapUsingMake["David"] = 10
	fmt.Println("Students Scores: ", studentsMapUsingMake)

	//deleting a key-value pair

	delete(studentsMapUsingMake, "Aayan")

	fmt.Println("Students Scores: ", studentsMapUsingMake)

	//check if key exists

	davidMarks, exists := studentsMapUsingMake["David"]

	if exists {
		fmt.Println("David Exists")
		fmt.Println("David Marks: ", davidMarks)
	} else {
		fmt.Println("David Doesn't Exists")
	}

	//Looping Through maps
	for name, marks := range studentsMapUsingMake {
		fmt.Printf("%s scored %d marks\n", name, marks)
	}
}
