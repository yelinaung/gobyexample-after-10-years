package main

import (
	"fmt"
)

// TODO Read https://go.dev/blog/slices-intro

func main() {
	var s []string

	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("after make:", s, "len:", len(s), "cap:", cap(s))
	fmt.Println(s == nil)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("after data assign:", s, "len:", len(s), "cap:", cap(s))

	// append requires to accept a return value, a slice
	s = append(s, "x")
	fmt.Println("append:", s)

	// we can also copy slice
	s2 := make([]string, len(s))
	copy(s2, s)
	fmt.Println("copy:", s2)

	// slicing
	// slice[low:high]
	// last item
	// no negative index like Python
	fmt.Println("last item:", s2[len(s2)-1])

	// only the first item
	// either with slicing or just .. s[0]
	fmt.Println("first item:", s2[:1])

	l := s[2:5]
	fmt.Println("sl1:", l)

	// The length can change as you add or remove elements,
	// but the capacity remains constant until you explicitly
	// resize the slice using the append() function (which ma trigger a reallocation).

	// create a slice with an initial capacity of 3
	fmt.Println()
	my_slice := make([]int, 0, 3)

	for i := 1; i <= 5; i++ {
		// Append elements to the slice
		my_slice = append(my_slice, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(my_slice), cap(my_slice))
	}
	fmt.Println()

	// testing the resizing before and after
	// use the & operator to print the memory address of the slice before and after the append() operation
	fmt.Println()
	my_slice2 := make([]int, 0, 3)
	// The capacity of a slice is "doubled" whenever a reallocation occurs.
	// This doubling strategy is part of Go's slice implementation and
	// helps optimize memory usage and reduce the number of reallocations,
	// making append operations more efficient.
	for i := 1; i <= 20; i++ {
		// This actually does not work lol
		// Print the memory address before appending
		fmt.Printf("Before append - Address: %p, Length: %d, Capacity: %d\n", &my_slice2, len(my_slice2), cap(my_slice2))

		// Append elements to the slice
		my_slice2 = append(my_slice2, i)

		// Print the memory address after appending
		fmt.Printf("After  append - Address: %p, Length: %d, Capacity: %d\n", &my_slice2, len(my_slice2), cap(my_slice2))
		fmt.Println()
	}
}
