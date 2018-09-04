package main

import (
	"fmt"
	"time"
)

func main() {
	// 初始化
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker at:", t)
		}
	}()
	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	var ticker2 *time.Ticker = time.NewTicker(50 * time.Millisecond)
	tc := make(chan time.Time, 5)
	go func() {
		for t2 := range ticker2.C {
			tc <- t2
			fmt.Println("Ticker At:", t2)
		}
	}()
	// for i := 0; i < 100; i++ {
	// 	if e, ok := <-tc; ok {
	// 		fmt.Println("Tick channel:", e)
	// 	}
	// }
	time.Sleep(1024 * time.Millisecond)
	ticker2.Stop()
}
