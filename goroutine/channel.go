package main

import (
	"fmt"
)

func main() {
	var ch1 chan int   // 读写
	var ch2 chan<- int // 只写
	var ch3 <-chan int // 只读
	// make初始化
	ci := make(chan int)
	cj := make(chan int, 0)
	cf := make(chan *os.File, 100) // 缓冲长度2147493526,2147493522
	//12333,58273 250909 70606 51783 lcp0578.github.io
}
