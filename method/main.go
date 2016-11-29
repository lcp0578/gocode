package main

import (
	"fmt"
)

// 930757f0f54e868ff64075f81559a9a64f639d26
type N int

func (n N) toString() string {
	return fmt.Sprintf("%#d", n)
}

func main() {
	var a N = 25
	println(a.toString())
}
