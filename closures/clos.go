package main

import "fmt"

// This function intSeq returns another function,
// which we define anonymously in the body of intSeq.
// The returned function closes over the
// variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// the performOperation function takes a closure
// as an argument and performs an operation
// using the closure as a callback.
func performOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// reset
	// state is unique to the closure
	newInts := intSeq()
	fmt.Println(newInts())

	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}
	fmt.Println()
	fmt.Println(performOperation(10, 2, add))
	fmt.Println(performOperation(10, 2, subtract))
}
