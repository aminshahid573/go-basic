package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map
	// string inside the bracket is type of key
	//string outside is type of value
	m := make(map[string]string)

	//setting an element
	m["name"] = "golang"
	m["version"] = "1.18"

	//get an element
	fmt.Println(m)
	fmt.Println(m["name"])

	//accessing non existing key
	fmt.Println(m["invalid"])
	// IMP: if key doesn't exist, it returns the zero value of the data type (e.g., 0 for int, "" for string, false for bool).

	m2 := make(map[string]int)
	m2["a"] = 1
	m2["b"] = 2
	m2["c"] = 3

	fmt.Println(m2["d"])

	// length of map
	fmt.Println("Length is ", len(m2))

	//delete an element
	delete(m2, "a")
	fmt.Println(m2)

	//clear whole map
	clear(m2)
	fmt.Println(m2)

	//declaring map with another method
	m3 := map[string]int{"price": 1040, "phones": 2}
	fmt.Println(m3)

	//checking and getting values
	//when accessing value in map it return two values
	//first is the value of key and second is true or false (wheter the key exists or not)
	//best practice to access value in map
	v, ok := m3["price"]
	fmt.Println(v)
	if ok {
		fmt.Println("All okay")
	} else {
		fmt.Println("not okay")
	}

	map1 := map[string]int{"price": 1040, "phones": 2}

	map2 := map[string]int{"price": 1040, "phones": 2}

	fmt.Println("are both map same : ", maps.Equal(map1, map2))

	//range over map
	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
