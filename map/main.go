package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]string)
	mp["a"] = "A"
	mp["b"] = "B"
	mp["c"] = "C"
	v, ok := mp["b"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key b not exist")
	}
	v, ok = mp["d"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key d not exist")
	}
}
