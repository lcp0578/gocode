package main

import (
	"fmt"
)

func main() {
	// if else
	a := 2
	if a == 2 {
		fmt.Println("ok")
	}

	if b := 2; b == 2 {
		fmt.Println("ok")
	}

	if c := 3; c < 3 {
		fmt.Println("c < 3")
	} else {
		fmt.Println("c=", c)
	}
	// switch, 默认带有break,可使用fallthrough去掉
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("default")
	}
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("default")
	}
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}
	result := 3
	switch {
	case result < 0:
		fmt.Println("result < 0")
	case result > 0:
		fmt.Println("result > 0")
	default:
		fmt.Println("result = 0")
	}
	// for
	hw := "Hello World"
	for i, j := 0, len(hw); i < j; i++ {
		fmt.Println(string(hw[i]))
	}
	num := 0
	for num < 5 {
		fmt.Println(num)
		num++
	}
	as := 65
	for {
		if as > 68 {
			break
		} else {
			fmt.Println(string(as))
			as++
		}
	}
	// 嵌套循环
outer:
	for i, j := 0, 5; i < j; i++ {
		for m, n := 0, 5; m < n; m++ {
			if m > 3 {
				break
			}
			if m == 2 {
				continue
			}
			if i == 4 {
				break outer
			}
			if i == 2 {
				continue outer
			}
			fmt.Println("i=", i, ",j=", j, ",m=", m, ",n=", n)
		}
	}
}
