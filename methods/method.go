package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perim())

	fmt.Println()

	// Go automatically handles conversion between values
	// and pointers for method calls.
	// You may want to use a pointer receiver type
	// to avoid copying on method calls or
	// to allow the method to mutate the receiving struct.
	r_pointer := &r
	fmt.Println("area:", r_pointer.area())
	fmt.Println("perimeter:", r_pointer.perim())
}
