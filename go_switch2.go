package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print(time.Now(), "\n")

	for i := 0; i < 10; i++ {
		z := i % 2
		switch {
		case 0 == z:
			fmt.Print("余り0\n")
		case 1 == z:
			fmt.Print("余り1\n")
		default:
		}
	}
}
