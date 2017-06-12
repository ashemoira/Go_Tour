package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	// a: n-1
	// b: n
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
