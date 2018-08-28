package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) printName() {
	fmt.Println(s.Name)
}
func (s *Student) getAge() int {
	return s.Age
}
func main() {
	s := Student{Name: "lcpeng", Age: 25}
	rt := reflect.TypeOf(s)
	// 若是指针，则取指针指向元素
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	fmt.Println(rt.PkgPath())
	fmt.Println(rt.Name())
	fmt.Println(rt.NumField())
	fmt.Println(rt.NumMethod()) // 1.7为0
	for i, j := 0, rt.NumField(); i < j; i++ {
		fmt.Println(rt.Field(i).Name)
	}
	prt := reflect.PtrTo(rt)
	fmt.Println(prt.Name(), prt.NumMethod())
}
