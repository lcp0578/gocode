package main

import (
	"fmt"
	//"os"
	"time"
)

func main() {
	//var ch1 chan int   //读写
	//var ch2 chan<- int //只写
	//var ch3 <-chan int //只读

	ch4 := make(chan int) //无缓冲区channel，也称同步的channel
	//ch5 := make(chan int, 0)      //同上
	//ch6 := make(chan *os.File, 2) //缓冲区为2
	list := []int{32, 17, 16, 62, 45, 76, 42, 90}
	go func() {
		sortList(list...)
		for _, val := range list {
			fmt.Println(val)
		}
		ch4 <- 1024
	}()
	doSomething()
	<-ch4
}

func doSomething() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

}

func sortList(args ...int) {
	length := len(args)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if args[i] > args[j] {
				tmp := args[i]
				args[i] = args[j]
				args[j] = tmp
			}
		}
	}
}
