package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Aaa struct {
	A int
	B int
	C string
	E string
}

type Bbb struct {
	B int
	C string
	D string
	F string
}

/**
eflect针对struct提供的函数方法如下：

func (v Value) Elem() Value //返回指针或者interface包含的值
func (v Value) Field(i int) Value //返回struct的第i个field
func (v Value) FieldByIndex(index []int) Value //返回嵌套struct的成员
func (v Value) FieldByName(name string) Value //通过成员名称返回对应的成员
func (v Value) FieldByNameFunc(match func(string) bool) Value //只返回满足函数match的第一个field
 **/
// reflect
func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()
	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}

func main() {
	aaa := &Aaa{1, 2, "ccc", "eee"}
	bbb := &Bbb{}
	CopyStruct(aaa, bbb)
	fmt.Println(bbb)
	ajson, _ := json.Marshal(aaa)
	bbbb := new(Bbb)
	_ = json.Unmarshal(ajson, bbbb)
	fmt.Println(bbbb)
}
