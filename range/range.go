package main

import "fmt"

func main() {
	// iterating over slice/array
	nums := []int{2, 3, 4, 5}
	for index, num := range nums {
		fmt.Printf("index - %d: value - %d\n", index, num)
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// iterating over map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("key - %s, value - %d\n", k, v)
	}

	// iterating over Unicode code point
	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}
