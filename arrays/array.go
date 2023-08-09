package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("current array:", a)

	a[2] = 100
	fmt.Println("current array:", a)

	// retrieve
	fmt.Println("array index 2:", a[2])

	// length
	fmt.Println("length:", len(a))

	// one-line to declare and initialize
	daysOfWeek := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i := 0; i < len(daysOfWeek); i++ {
		fmt.Printf("%d, %s\n", i+1, daysOfWeek[i])
	}

	// 2D array
	var twoD [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", twoD[i][j])
		}
		fmt.Println()
	}
}
