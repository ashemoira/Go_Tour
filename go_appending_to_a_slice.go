package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)

	// We can add more then one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

	var m []int
	printSlice(m)

	for i := 0; i < 50; i++ {
		m = append(m, 1)
		printSlice(m)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
