package main

import (
	"fmt"
)

func main() {
	const PI = 3.1415
	const HI = "Hello world"
	const (
		z   = false
		abc = 123
	)
	const username, truename = "lcp0578", "lcpeng"
	fmt.Println("PI=", PI, ",HI=", HI, ",z=", z, ",abc=", abc, ",username=", username, ",truename=", truename)

	//模拟枚举类型
	const (
		Sunday = iota
		Monday
		Tuesday
	)
	fmt.Println("Sunday=", Sunday, ",Monday=", Monday, ",Tuesday=", Tuesday)
	//模拟枚举类型
	const (
		a, A = iota, iota
		b, B
		c, C
	)
	fmt.Println("a=", a, ",b=", b, ",c=", c, ",A=", A, ",B=", B, ",C=", C)
	//模拟枚举类型
	const (
		Sun = iota
		Mon
		Tues
		Wed  = "叁"
		Thur = iota
		Fri
	)
	fmt.Println("Sun=", Sun, ",Mon=", Mon, ",Tues=", Tues, ", Wed=", Wed, ", Thur=", Thur, ", Fri=", Fri)
}
