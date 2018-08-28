package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		//quit <- 1
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
