package main

import "fmt"

type numbers []int

func newNumbers() numbers {
	n := numbers{}

	for i := 0; i <= 10; i++ {
		n = append(n, i)
	}

	return n
}

func (n numbers) decision () {
	for _, v := range n {
		if v % 2 == 0 {
			fmt.Printf("%d is even\n", v)
		} else {
			fmt.Printf("%d is odd\n", v)
		}
	}
}
