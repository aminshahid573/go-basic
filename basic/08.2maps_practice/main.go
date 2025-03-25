package main

import "fmt"

func main() {
	country := map[string]int{
		"India": 100000,
		"USA":   20000,
		"UK":    3000,
		"Nepal": 40000,
	}

	country["Bhutan"] = 50000

	country["Nepal"] = 32501

	delete(country, "India")

	for country, population := range country {
		fmt.Printf("%s has %d Population\n", country, population)
	}
}
