package main

import (
	"fmt"
)

type Student struct {
	id       int
	username string
	age      int
}

func (this *Student) getUsername() string {
	return this.username
}
func (this *Student) getAge() int {
	return this.age
}

type IsStudent interface {
	getUsername() string
	getAge() int
}

func main() {
	var s1 IsStudent = &Student{1, "lcpeng", 25}
	fmt.Println(s1)
}
