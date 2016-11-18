package main

import (
	"errors"
	"fmt"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func log(x int, err error) {
	fmt.Println(x, err)
}

func test(x, y int) (int, error) {
	return div(x, y) // 多返回值作为return结果
}

func main() {
	x := 1024
	y := 0
	log(test(x, y)) // 多返回值作实参
}
