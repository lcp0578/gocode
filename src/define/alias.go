package main

import (
	"fmt"
)

func main() {
	// byte alias for uint8
	// rune alias for int32
	var a byte = 0x11
	var b uint8 = a
	test(a)
	test(b)
	var c uint8 = a + b
	test(c)
	var d bool = a == b
	fmt.Println(d)

	var m rune = 1024
	var n int32 = m
	println(n)
	//var x int = 100
	//var y int64 = x
	//add(x, y)

}

func test(x byte) {
	println(x)
	fmt.Println(x)
}
