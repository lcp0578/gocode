package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 100
	fmt.Println(&x)

	//x, y := 200, "lcpeng" //退化赋值操作
	//fmt.Println(&x, x)
	//fmt.Println(y)
	{
		x, z := 200, "lcpeng.com" //退化赋值操作
		fmt.Println(&x, x)
		fmt.Println(z)
	}

	m, n := 1, 2
	m, n = m+n, m+n
	fmt.Println(m, n)

	s, _ := strconv.Atoi("lcp") // _垃圾桶
	println(s)
	ss, err := strconv.Atoi("lcp") // _垃圾桶
	if err != nil {
		println(ss)
	} else {
		err.Error()
	}

}
