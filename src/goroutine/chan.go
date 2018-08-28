package main

import (
	"fmt"
)

func producer(c chan int) {
	defer close(c) // close channel
	for i := 0; i < 10; i++ {
		c <- i // 阻塞
	}
}

func consumer(c, f chan int) {
	for {
		if v, ok := <-c; ok {
			fmt.Println(v) // 阻塞
		} else {
			fmt.Println(ok)
			break
		}
	}
	f <- 1
}

func main() {
	buf := make(chan int)
	flg := make(chan int)
	go producer(buf)
	go consumer(buf, flg)
	<-flg
}
