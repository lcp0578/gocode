package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S //匿名嵌入字段
}

func (S) sVal() {}

func (*S) sPtr() {}

func (T) tVal() {}

func (*T) tPtr() {}

// 显示方法集里所有方法名字
func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
func main() {
	var s S
	var t T

	methodSet(s)
	println("========")
	methodSet(&s)
	println("----------")

	methodSet(t)
	println("========")
	methodSet(&t)
	println("----------")

}
