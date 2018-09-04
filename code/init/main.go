package main

import (
	"fmt"
)

func main() {
	type data struct {
		x int
		s string
	}
	var a data = data{1024, "lcpeng"}
	b := data{
		2048,
		"lcp0578",
	}
	c := []int{
		1,
		2}
	d := []int{1, 2, 3,
		4,
		5,
	}
	fmt.Println(a, b, c, d)
}
