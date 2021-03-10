package main

import "fmt"

func main() {
	var s = []int{2, 3, 5, 6, 2, 2, 3, 4, 4, 32, 88, 3, 7, 8, 4, 3, 45, 23, 44, 7, 5, 3, 0, 99, 1, 1, 1}

	for _, e := range s {
		fmt.Printf("Mul: %d\n", e*2)
	}

	for _, e := range s {
		fmt.Printf("Add: %d\n", e+2)
	}

	for _, e := range s {
		fmt.Printf("Div: %d\n", e/2)
	}

	for _, e := range s {
		fmt.Printf("Sub: %d\n", e-2)
	}
}
