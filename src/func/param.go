package main

import (
	"fmt"
)

// 值拷贝传递 pass-by-value
func test(x *int) {
	fmt.Printf("pointer: %p, target %v\n", &x, x)
}

func main() {
	a := 0x100
	p := &a
	fmt.Printf("Pointer: %p, Target %v\n", &p, p)
	test(p)
	x := 100
	p2 := &x
	test2(p2)

	var p3 *int
	test3(&p3)
	println(*p3)
}

func test2(p *int) {
	go func() {
		println(p)
	}()
}

func test3(p **int) {
	x := 100
	*p = &x
}
