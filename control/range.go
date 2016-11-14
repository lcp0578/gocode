package main

import (
	"fmt"
)

func main() {
	data := [3]string{"l", "c", "p"}
	for i, s := range data {
		fmt.Println(i, "=>", s)
		fmt.Println(&i, &s) // 定义的局部变量会重复使用
	}
	for i := range data {
		fmt.Println(i, data[i])
	}
	for _, s := range data {
		fmt.Println(s)
	}

	num := [3]int{10, 20, 30}
	for i, x := range num { // 从data的复制品取值
		if i == 0 {
			num[0] += 100
			num[1] += 200
			num[2] += 300
		}
		fmt.Println("x=", x, ",num[", i, "]", num[i])
	}
	num2 := [3]int{10, 20, 30}
	for i, x := range num2[:] { // i == 0时，x已取值，所以x = 10
		if i == 0 {
			num2[0] += 100
			num2[1] += 200
			num2[2] += 300
		}
		fmt.Println("x=", x, ",num2[", i, "]", num2[i])
	}
}
