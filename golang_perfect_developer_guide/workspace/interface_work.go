package main

import "fmt"

type bird interface {
	chirp() string
	fly() string
}

type fether struct {}
type stone struct {}

func main() {
	fe := fether{}
	st := stone{}

	printChirp(fe)
	printChirp(st)
}

func printChirp(b bird) {
	fmt.Println(b.chirp())
}

func printFly(b bird) {
	fmt.Println(b.fly())
}

func (f fether) chirp() string {
	return "fether"
}

func (s stone) chirp() string {
	return "stone"
}

func (f fether) fly() string {
	return "fly"
}

func (s stone) fly() string {
	return "don't fly ..."
}
