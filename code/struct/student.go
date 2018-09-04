package main

import (
	"fmt"
	"learn-go/student"
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
	cs := CollegeStudent{st2, "计算机科学与技术"}
	fmt.Println(cs)
	fmt.Println(cs.Student.username)
	fmt.Println(cs.username)
	fmt.Println(cs.getName())
	mys := student.MyStudent{1, "大鹏", 2}
	fmt.Println(mys)
}

func (this Student) getName() string {
	return this.username
}

func (this *Student) getAge() int {
	return this.age
}

type CollegeStudent struct {
	Student
	Profession string
}

func (this CollegeStudent) getName() string {
	return this.username + this.Profession
}
