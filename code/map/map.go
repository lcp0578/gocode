package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1024
	x, ok := m["b"] // ok-idiom模式
	fmt.Println(x, ok)
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
}
