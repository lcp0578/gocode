package main

import (
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt("10110011", 2, 32)
	b, _ := strconv.ParseFloat("01273", 64)
	c, _ := strconv.ParseInt("64", 16, 32)
	d, _ := strconv.ParseInt("64", 16, 64)
	println(a)
	println(b)
	println(c)
	println(d)
}
