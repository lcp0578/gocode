package main

import (
	"fmt"
)

func main() {
	r1, r2 := test(5, 2)
	fmt.Println("add=", r1, ",sub=", r2)
	s1, s2 := test2(5, 2)
	fmt.Println("add=", s1, ",sub=", s2)
	result := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
	slice := []int{1, 2, 3, 4, 5, 6}
	result2 := sum(slice...)
	fmt.Println(result2)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func test(a, b int) (int, int) {
	return add(a, b), sub(a, b)
}

// 命名返回值
func test2(a, b int) (addRst, subRst int) {
	addRst = add(a, b)
	subRst = sub(a, b)
	return
}

func sum(args ...int) int {
	s := 0
	for _, val := range args {
		s += val
	}
	return s
}
