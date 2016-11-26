package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	p := &s
	p0 := &s[0]
	p1 := &s[1]

	fmt.Println(p, p0, p1)
	(*p)[0] += 100
	*p1 += 100
	fmt.Println(*p, s)

	x := [][]int{
		{1, 2},
		{4, 5, 6},
		{7},
	}
	fmt.Println(x[1])
	x[2] = append(x[2], 8, 9)
	fmt.Println(x[2])
}
