package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := [...]int{2, 3, 5, 7, 11, 13}
	s1 := p[1:3]
	fmt.Println(p)
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println("changeArrayValue")
	changeArrayValue(p)
	fmt.Println(p)
	fmt.Println(s1)
	fmt.Println("changeSliceValue")
	changeSliceValue(s1)
	fmt.Println(p)
	fmt.Println(s1)
	fmt.Println("changeSliceValue2")
	changeSliceValue2(s1)
	fmt.Println(p)
	fmt.Println(s1)
	arr := [4]int{3: 1}
	s2 := arr[1:1]
	fmt.Println(arr)
	fmt.Println(s2)
	s3 := arr[3:4]
	fmt.Println(s3)
	var hel = "hello"
	shel := []byte(hel)
	shel[0] = 'c'
	strhel := string(shel)
	fmt.Println(shel, strhel)
}

func changeArrayValue(arr [6]int) {
	arr[0] = 100
}

func changeSliceValue(slice []int) {
	slice[0] = 100
}

func changeSliceValue2(slice []int) {
	slice[1] = 100
}
