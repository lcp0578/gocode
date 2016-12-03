package main

import (
	"fmt"
)

type N int

func (n N) test() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func main() {
	var n N = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1 := N.test
	f1(n)

	f2 := (*N).test
	f2(&n)
	fmt.Println("==========")
	N.test(n)
	(*N).test(&n)
}
