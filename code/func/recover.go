package main

import (
	"fmt"
)

func main() {
	test(5, 0)
}

func divide(n, m int) int {
	return n / m
}

func test(n, m int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	s := divide(n, m)
	panic("error")
	return s
}
