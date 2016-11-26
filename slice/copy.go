package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := s[5:8]
	s2 := s[4:]
	fmt.Println(s, s1, s2)
	n := copy(s2, s1) // 最终复制长度以较短切片长度为准
	fmt.Println(s, s1, s2, n)

	ss := make([]int, 6)
	fmt.Println(ss)
	n = copy(ss, s)
	fmt.Println(s, ss, n)

	// 从字符串中复制数据到[]byte
	b := make([]byte, 3)
	n = copy(b, "abcdef")
	fmt.Println(b, n)
}
