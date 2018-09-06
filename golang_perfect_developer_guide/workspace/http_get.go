package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	body, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up!")
	fmt.Println(body)
}
