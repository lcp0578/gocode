package main

import (
	"fmt"
)

type N int

func (n N) testVal() {
	fmt.Printf("testVal.n: %p, %v\n", &n, n)
}

func (n *N) testPtr() {
	fmt.Printf("testPtr.n: %p, %v\n", n, *n)
}
func main() {
	var n N = 100
	p := &n

	n++
	fv1 := n.testVal // 复制
	fp1 := n.testPtr // 复制指针

	n++
	fv2 := p.testVal //复制
	fp2 := p.testPtr // 复制指针

	n++
	fmt.Printf("main.n: %p, %d\n", p, n)
	fv1()
	fp1()
	fv2()
	fp2()
}
