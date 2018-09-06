package main

import (
	"os"
	"fmt"
	"bufio"
)

const BUFSIZE = 1024

func main() {
}

func file_create() {
	file, err := os.Create(`tmp_list`)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	output := "todo list"
	file.Write(([]byte)(output))
}

func file_open() {
	file, err := os.Open(`tmp_list`)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		err := sc.Err()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
