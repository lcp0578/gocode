package main

import (
	"fmt"
)

func main() {
	const v = 20

	var a byte = 10
	b := v + a // v自动转换为byte/unit8类型
	fmt.Printf("%T, %v\n", b, b)
	const c float32 = 1.2
	d := c + v // v 自动转换为float32
	fmt.Printf("%T, %v\n", d, d)
}
