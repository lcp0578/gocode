package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	Name string
	Age  int
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
}
