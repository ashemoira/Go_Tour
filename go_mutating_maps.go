package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Result"] = 100
	v, ok = m["Result"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Set"] = 50
	v, ok = m["Set"]
	fmt.Println("The value:", v, "Present?", ok)

	// all output
	fmt.Println(m)
}
