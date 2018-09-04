package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

func main() {
	u := map[int]user{
		1: {"Tom", 19},
	}
	//u[1].age += 1  // cannot assign to struct field u[1].age in map
	u1 := u[1]
	u1.age += 1
	fmt.Println(u1)

	uu := map[int]*user{
		1: &user{"Jack", 20},
	}
	uu[1].age++ //uu[1]返回的是指针
	fmt.Println(uu[1])
}
