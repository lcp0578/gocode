package main

import (
	"fmt"
)

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := x[:]     //len=10,cap=10
	s2 := x[2:5:7] //len=3,cap=7-2
	s3 := x[2:5]   //len=3,cap=len(x) - 2
	s4 := x[:4]    //len=4,cap=len(x) - 0
	s5 := x[4:]    //len=6,cap=len(x) - 4
	s6 := x[:4:6]  //len=4,cap=6
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s4), cap(s4))
	fmt.Println(len(s5), cap(s5))
	fmt.Println(len(s6), cap(s6))
}
