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

func div2(x, y int) (r int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
		return
	}
	r = x / y
	return
}
func div3(x, y int) (r int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
		return
	}
	{
		r := x / y // 新定义的同名局部变量，同名遮蔽
		return r, err
	}
	return
}
func log(x int, err error) {
	fmt.Println(x, err)
}

func test(x, y int) (int, error) {
	//return div(x, y) // 多返回值作为return结果
	//return div2(x, y)
	return div3(x, y)
}

func main() {
	x := 1024
	y := 0
	log(test(x, y)) // 多返回值作实参
}
