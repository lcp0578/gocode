package main

import (
	"fmt"
)

func test(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	x := 1024
	f := test(x)
	f()
}
