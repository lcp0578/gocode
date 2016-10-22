package main

import (
	"fmt"
)

// struct 成员大写字母开头才可以在包外访问
type Student struct {
	id       int
	username string
	sex      string
	age      int
}

func main() {
	st1 := new(Student)
	st1.id = 2
	st1.username = "lcp0578"
	st1.sex = "男"
	st1.age = 25
	fmt.Println(st1)
	st2 := Student{3, "lcpeng", "女", 26}
	fmt.Println(st2)
	st3 := Student{id: 4, username: "dapeng", sex: "nan", age: 18}
	fmt.Println(st3.getName(), st3.getAge())
}

func (this Student) getName() string {
	return this.username
}

func (this *Student) getAge() int {
	return this.age
}
