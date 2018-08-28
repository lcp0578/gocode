package main

import (
	"fmt"
	"time"
)

func main() {
	// 初始化定时器
	nt := time.NewTimer(2 * time.Second)
	// 当前时间
	now := time.Now()
	fmt.Println("Now time:", now)
	expire := <-nt.C
	fmt.Println("Expire time:", expire)

	//初始化通道
	ch1 := make(chan int, 500)
	sign := make(chan byte, 1)
	// 给ch1通道写入数据
	for i := 0; i < 500; i++ {
		ch1 <- i
	}
	// 单独起一个Goroutine执行select
	go func() {
		var e int
		ok := true
		var timer *time.Timer
		for {
			select {
			case e = <-ch1:
				fmt.Printf("ch1 -> %d \n", e)
			case <-func() <-chan time.Time {
				if timer == nil {
					// 初始化到期时间据此间隔1ms的定时器
					timer = time.NewTimer(time.Millisecond)
				} else {
					// 复用，通过Reset方法重置定时器
					timer.Reset(time.Millisecond)
				}
				// 得知定时器事件来临时，返回结果
				return timer.C
			}():
				fmt.Println("Timeout")
				ok = false
				break
			}
			// 终止for循环
			if !ok {
				sign <- 0
				break
			}
		}
	}()
	// 惯用手法，读取sing的数据，为了等待select的Goroutine执行
	<-sign
}
