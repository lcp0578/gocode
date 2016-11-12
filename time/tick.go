package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	c := time.After(3 * time.Second) // 返回channel类型，3s后发送当前时间
	t := <-c
	fmt.Println(t)
	nt := time.NewTimer(3 * time.Second) //返回Timer类型
	ntt := nt.C
	fmt.Println(ntt)
	var tc chan int
	time.AfterFunc(5*time.Second, Test(tc))
	<-tc
	//time.Sleep(8 * time.Second)
	// var str string
	// fmt.Scan(&str)

	// tk := time.Tick(2 * time.Second)
	// for tkt := range tk {
	// 	fmt.Println(tkt)
	// }
}

func Test(ct chan int) {
	ct <- 1
	fmt.Println("hello golang", time.Now())
}
