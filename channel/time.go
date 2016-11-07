package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("input data")
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}
