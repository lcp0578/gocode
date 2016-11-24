package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

type CollageStudent struct {
	XMLName xml.Name `xml:"collegeStudent"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Phones  []string `xml:"phones>phone"`
}

func main() {
	str := `<?xml version="1.0" encoding="utf-8" ?>
			<Student>
				<Name>李四</Name>
				<Age>25</Age>
			</Student>`

	var s Student
	xml.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	cstr := `<?xml version="1.0" encoding="utf-8" ?>
			<collegeStudent>
				<name>李四</name>
				<age>25</age>
				<phones>
					<phone>2048</phone>
					<phone>1024</phone>
				</phones>
			</collegeStudent>`
	var cs CollageStudent
	xml.Unmarshal([]byte(cstr), &cs)
	fmt.Println(cs)
}
