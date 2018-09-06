package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Print("started.")

	sleep5_finished := make(chan string)

	go waitFiveSec(sleep5_finished)
	fmt.Println(<- sleep5_finished)

	log.Print("all finished.")
}

func waitFiveSec(sleep5_finished chan string) {
	log.Print("sleep5 start.")
	time.Sleep(5 * time.Second)
	log.Print("sleep5 finished.")
	sleep5_finished <- "end time"
}
