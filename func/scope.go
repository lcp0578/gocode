package main

import (
	"fmt"
)

func main() {
	var result = f()
	var result2 = f2()
	fmt.Println(result, result2, "end")
}

func f() (result int) {
	fmt.Println(result)
	defer func() {
		fmt.Println(result)
		result++
		fmt.Println(result)
	}()
	return 1
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
		fmt.Println(t)
	}()
	return t
}
