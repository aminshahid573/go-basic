package main

import (
	"fmt"
	"slices"
)

// slice is dyanameic array
// most used contruct
func main() {
	//uninitialized slice is nil
	var nums []int
	fmt.Println(nums)
	if nums == nil {
		fmt.Println("Slice is Empty")
	}

	//to get length of slice
	fmt.Println(len(nums))
	/*
		   	🛠️ What is make in Go?
		   In Go, make is a built-in function used to create slices, maps, and channels with non-zero initialization. Unlike new, which returns a pointer to a zeroed value, make returns the actual object, already initialized and ready to use.
		   make(type, length)
			make(type, length, capacity) // For slices

	*/

	// make(type, length, capacity)
	var nums2 = make([]int, 2)
	// Length (len): The number of elements the slice currently holds.
	//Capacity (cap): The number of elements the slice can hold before needing to allocate a bigger array.(kindof maximum but can grow)
	fmt.Println((nums2))      // [0 0]
	fmt.Println(nums2 == nil) //false (not nil)

	//capacity
	fmt.Println("Capacity is : ", cap(nums2)) // capacity is 2

	//append value
	nums2 = append(nums2, 1)
	fmt.Println(nums2)
	fmt.Println("Capacity is : ", cap(nums2)) // capacity is 4 (unexpected) explanation below

	/*
			You're appending a new element to a slice that already has full capacity (len == cap == 2)

		Go must create a new underlying array to hold the new size (3)

		Go's internal strategy is to double the capacity for efficiency → So new capacity becomes 4

		The original contents [0 0] are copied into the new array, and 1 is appended
	*/

	//recommened way to create slice
	// var nums = make([]int,0,estimatedsize)
	// you can alos do like
	// nums := []int{}

	//accessing and changing values
	nums3 := make([]int, 2)

	nums3[0] = 10
	fmt.Println(nums3)
	fmt.Println(nums3[0])

	// copy the slice
	nums4 := make([]int, len(nums3))
	fmt.Println("before copy ", nums3, nums4)

	copy(nums4, nums3)
	fmt.Println("After copy ", nums3, nums4)

	//slice operator
	var nums5 = []int{1, 2, 3}
	fmt.Println(nums5[0:2]) //start with 0th and upto 2th and exclude 2th value
	// we can also use it like nums5[:1] it is ame as nums5[0:1]
	// we can also use it like nums5[1:] it return value from 1th to all in right

	var nums6 = []int{1, 2}
	var nums7 = []int{1, 2}
	fmt.Println(slices.Equal(nums6, nums7))

	//multidimesional slices
	var nums8 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums8)
}
