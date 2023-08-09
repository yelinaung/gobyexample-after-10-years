package main

import "fmt"

// Go version of dict/HashMap
// https://go.dev/blog/maps
func main() {
	m := make(map[string]int)
	m["age"] = 12
	m["k1"] = 11
	m["k2"] = 22
	fmt.Println(m)

	// retrieve
	fmt.Println(m["k2"])

	// retrieve non-existent key
	// retruns a zero value
	// https://go.dev/ref/spec#The_zero_value
	fmt.Println(m["k3"])

	m2 := make(map[string]string)
	m2["name"] = "yelinaung"
	m2["city"] = "bangkok"
	fmt.Println(m2["k2"])

	delete(m2, "name")
	delete(m2, "something")
	fmt.Println(m2)

	n := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(n)
}
