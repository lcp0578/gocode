package main

import "fmt"

type bindFunc func(int, int) int

type loudBindFunc func(int, int) int

func add(x, y int) int {
	return x + y
}

//在go中任何类型具有Error() string方法就是有效的error类型
func (f bindFunc) Error() string {
	return "bindFunx error"
}

func (f loudBindFunc) Error() string {
	return "THIS IS LOUD BIND FUNC ERROR"
}

func main() {
	var err error
	err = bindFunc(add)
	fmt.Println(err.Error())
	err = loudBindFunc(add)
	fmt.Println(err.Error())
}
