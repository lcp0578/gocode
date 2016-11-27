package main

import (
	"fmt"
	"unsafe"
)

type point struct {
	x, y int
}

type value struct {
	id    int    //基本类型
	name  string // 字符串
	data  []byte // 引用类型
	next  *value // 指针类型
	point        //匿名字段
}

func main() {
	v := value{
		id:    1,
		name:  "lcpeng",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200},
	}

	s := `
		v:%p ~ %x, size: %d, align: %d

		field    address    offset   size
		-------|----------|---------|-------
		id 		 %p 		%d 		  %d
		name 	 %p 		%d 		  %d
		data 	 %p 		%d 		  %d
		next	 %p 		%d 		  %d
		x 		 %p 		%d 		  %d
		y 		 %p 		%d 		  %d
	`
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y),
	)
}
