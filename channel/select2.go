package main 

import (
	"time"
	"fmt"
	)
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1);
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	forTag:
	for {
		select {
			case msg1 := <-c1:
				fmt.Println("receiveed", msg1)
			case msg2 := <-c2:
				fmt.Println("receiveed", msg2) 
				break forTag
		}
	}


	for {
		select {
			case msg1 := <-c1:
				fmt.Println("receiveed", msg1)
			case msg2 := <-c2:
				fmt.Println("receiveed", msg2) 
				goto gotoTag
		}
	}
	gotoTag:

	for {
		select {
			case msg1 := <-c1:
				fmt.Println("receiveed", msg1)
			case msg2 := <-c2:
				fmt.Println("receiveed", msg2) 
				return 
		}
	}


	for i := 0; i < 2; i++ {
		select {
			case msg1 := <-c1:
				fmt.Println("receiveed", msg1)
			case msg2 := <-c2:
				fmt.Println("receiveed", msg2) 
		}
	}
}