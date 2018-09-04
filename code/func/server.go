package main

import (
	"fmt"
	"log"
	"time"
)

type serverOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func newOption() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func server(option *serverOption) {
	fmt.Println(*option)
}

// 变参
func test(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a)
}

func main() {
	opt := newOption()
	opt.port = 9001
	server(opt)
	test("abc", 1, 2, 3, 4, 5)
	arr := [3]int{10, 20, 30}
	test("def", arr[:]...)
}
