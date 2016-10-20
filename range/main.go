package main

import (
	"fmt"
)

func main() {
	arr := [3]int{2, 4, 6}
	var mp = map[int]string{1: "A", 2: "B", 3: "C"}
	for _, v := range arr {
		fmt.Println(v)
	}
	for k, v := range mp {
		fmt.Println(k, "=", v)
	}
}
