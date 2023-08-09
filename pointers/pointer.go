package main

import "fmt"

// argument passing in by 'value'
// zeroval gets a copy of the argument, ival
func zeroval(ival int) {
	ival = 0
}

// Pass by Reference
// argument is int pointer
// *iptr then dereferences the pointer
// from its memory address
// Assigning a value to a dereferenced pointer
// changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

// pointer is a special variable
// that holds the memory address of another variable
// https://go.dev/ref/spec#Pointer_types
func main() {
	a := 42

	// &a syntax gives the memory address of 'a'
	// which is a pointer to 'a'
	ptr_a := &a
	fmt.Println("a :", a)
	fmt.Println("ptr_a :", ptr_a)
	fmt.Println("value of 'a' through pointer :", *ptr_a)

	// modifying the value of x through the pointer ptr.
	// The changes made to *ptr directly affect
	// the original variable x.
	*ptr_a = 100
	fmt.Println("after modification a :", a)

	fmt.Println()

	// reassigning
	x := 20
	fmt.Println("before modification x :", x)
	x = 100
	fmt.Println("after modification x :", x)

	fmt.Println()

	i := 70
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("after zeroval:", i)

	zeroptr(&i)
	fmt.Println("after zeroptr:", i)

	fmt.Println("after everything:", i)

	// printing the pointer
	fmt.Println("after zeroptr:", &i)
}
