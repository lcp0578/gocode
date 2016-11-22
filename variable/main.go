package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i32 int32 = 0xfff
	var i64 int64 = 0xfff
	fmt.Println(unsafe.Sizeof(i32), unsafe.Sizeof(i64))
	var f32 float32 = 1.23
	var f64 float64 = 1.23
	fmt.Println(unsafe.Sizeof(f32), unsafe.Sizeof(f64))
}
