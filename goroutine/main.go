package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	k := runtime.GOMAXPROCS(runtime.NumCPU())
	go SayHello()
	go SayWorld()
	go TestExit()
	time.Sleep(5 * time.Second)
	fmt.Println(k)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
		fmt.Println("Hello", runtime.NumGoroutine())
		runtime.Gosched()
	}
}
func SayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("World")
		fmt.Println("World", runtime.NumGoroutine())
		runtime.Gosched()
	}
}

func TestExit() {
	defer func() {
		fmt.Println("in defer")
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for i > 5 {
			runtime.Goexit()
		}
	}
}
