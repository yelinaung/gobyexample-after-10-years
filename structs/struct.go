package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	bob := person{"Bob", 20}
	fmt.Println(bob)

	alice := person{name: "Alice", age: 30}
	fmt.Println(alice)

	josh := person{name: "Josh"}
	fmt.Println(josh)

	// An '&' prefix yields a pointer to the struct.
	fred := &person{name: "Fred", age: 25}
	fmt.Println(fred)

	john := newPerson("John", 30)
	fmt.Println(john)
	fmt.Println(john.age)
	fmt.Println(john.name)

	fmt.Println()
	sean := person{name: "Sean", age: 50}
	fmt.Println(sean.name)
	sean_pointer := &sean
	fmt.Println(sean_pointer)
	fmt.Println(sean_pointer.name)

	sean.age = 51
	fmt.Println(sean.age)
	fmt.Println(sean)

	sean_pointer.age = 49
	fmt.Println(sean)
}
