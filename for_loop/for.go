package main

import "fmt"

func main() {
	var i int = 1
	for i <= 3 {
		fmt.Println(i + 1)
	}

	for {
		fmt.Println("Loop")
		break
	}

	// continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
