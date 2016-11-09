package main

import (
	"fmt"
)

const LCP int = 10 // 可以不使用

const (
	size int64 = 1024
	eof        = -1 //int
)

func main() {
	fmt.Println("hello const")
	const PI float64 = 3.1416926
	//fmt.Println(&PI, PI)
	fmt.Println(PI)

	const UA string = "chrome"
	//fmt.Println(&UA, UA)
	// 预定义常量, true、false、iota
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	const (
		c4 = iota
		c5 = iota
		c6 = iota
	)
	const (
		c7 = 1 << iota
		c8 = 1 << iota
		c9 = 1 << iota
	)
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9)
	const (
		con1 = iota
		con2
		con3
	)
	const (
		con4 = iota
		con5
		con6
	)
	const (
		con7 = 1 << iota
		con8
		con9
	)
	fmt.Println(con1, con2, con3, con4, con5, con6, con7, con8, con9)
}
