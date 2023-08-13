package main

import "fmt"

// MapKeys takes a map of any type and returns a slice of its keys.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func main() {
	fmt.Println("=== Generics ===")
	m := map[int]string{1: "one", 2: "two"}
	fmt.Println(MapKeys(m))

	m2 := map[int]float64{1: 1.1, 2: 2.2}
	fmt.Println(MapKeys(m2))
}
