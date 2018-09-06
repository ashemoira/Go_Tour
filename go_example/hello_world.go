package main

import "fmt"

type greeter interface {
	greet()
}

type japanese struct{}
type american struct{}

func (ja japanese) greet() {
	fmt.Println("こんにちは")
}

func (us american) greet() {
	fmt.Println("hello")
}

func main() {
	greeters := []greeter{japanese{}, american{}}

	for _, greeter := range greeters {
		greeter.greet()
	}
	japanese{}.greet()
}
