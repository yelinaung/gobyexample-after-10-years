package main

import "fmt"

// This is similar to Python *args
func sum_of_ints(nums ...int) int {
	fmt.Println("nums:", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum_of_ints(1, 2, 3, 4, 5))
	fmt.Println(sum_of_ints(2, 3, 4))
	fmt.Println(sum_of_ints())
}
