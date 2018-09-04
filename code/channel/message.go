package main

import (
	"fmt"
)

func main() {
	queue := make(chan int, 1)
	go producer(queue)
	go consumer(queue)
}

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
	}
}

func consumer(queue <-chan int) {
	for v := <-chan{
		fmt.Println("Receive", v)
	}
}
