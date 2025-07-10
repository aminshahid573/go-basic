package main

import "fmt"

func main() {

	//range in slice
	nums := []int{6, 7, 8}

	// //using classing for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//i is index and num is value
	for i, num := range nums {
		fmt.Println(i, num)
	}

	//range in map
	m := map[string]string{"fname": "Shahid", "lname": "Amin"}

	for key, value := range m {
		fmt.Println(key, value)
	}

	//range in string
	//give starting byte of rune (i) and unicode point rune (c)
	// i is not index
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}
