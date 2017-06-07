package main

import "fmt"

func main() {
	sum := 1
	for sum < 5000 {
		fmt.Print(sum, "\n")
		sum += sum
	}
	fmt.Println(sum)
}
