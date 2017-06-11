package main

import "fmt"

type Vertex struct {
	Lab, Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex, 5)
	// set maps value
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967, 54,
	}
	m["Bell"] = Vertex{
		1, 2, 3,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Bell"])
	fmt.Println(len(m))
}
