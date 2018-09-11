package main

import (
	"fmt"
	"log"
	"runtime"
)

func worker(msg string) <-chan string {
	// bufferを5に指定
	limit := make(chan int, 5)
	receiver := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			// NumGoroutineは現在起動しているgoroutineの数を知ることができる
			// 今回はbufferの数を5としているのでroutineは最大5を期待しているが、
			// 他のroutineが動いていることもあるので必ずしも上限が5とは限らない
			log.Println(runtime.NumGoroutine())
			// limitに送信する
			// 今回はbufferを5に指定しているので、5回送信がされた時点でchannelがブロックされる
			// 受信が完了したらまた動かすことができる
			// 送信: limit <- i
			// 受信: <-limit
			limit <- i
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, msg)
				receiver <- msg
				log.Println("limit :" , <-limit)
			}(i)
		}
	}()
	return receiver
}

func main() {
	receiver := worker("job")
	for i := 0; i<100 ; i++ {
		log.Println(<-receiver)
	}
}
