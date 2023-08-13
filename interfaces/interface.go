package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go,
// we just need to implement all the methods in the interface.
// Here we implement geometry on rects.
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area :", g.area())
	fmt.Println("perimeter:", g.perimeter())
}

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type Duck struct {
	name string
}

type JavaProgrammer struct {
	name string
}

// implement the Speak method for Dog
func (d Dog) Speak() string {
	return "Woof ! My name is " + d.name
}

func (c *Cat) Speak() string {
	return "Meow! My name is " + c.name
}

func (d Duck) Speak() string {
	return "Quack! My name is " + d.name
}

func (j JavaProgrammer) Speak() string {
	return "Hello World! My name is " + j.name
}

func speakNow(a Animal) {
	fmt.Println(a)
	fmt.Println(a.Speak())
}

func main() {
	r := rectangle{width: 3, height: 4}
	measure(r)

	c := circle{radius: 5}
	measure(c)

	fmt.Println()

	animals := []Animal{&Dog{"Max"}, &Cat{"Leo"}, Duck{"Donald"}, JavaProgrammer{"John"}}
	for _, animal := range animals {
		speakNow(animal)
	}
}
