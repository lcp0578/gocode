package main

import (
	"fmt"
)

type MyFuncType func(int) bool

func IsBigThan5(n int) bool {
	return n > 5
}
func IsSmallThan5(n int) bool {
	return n < 5
}
func Display(arr []int, f MyFuncType) {
	for _, val := range arr {
		if f(val) {
			fmt.Println(val)
		}
	}
}
func main() {
	arr := []int{3, 7, 1, 12}
	Display(arr, IsBigThan5)
	Display(arr, IsSmallThan5)
}
