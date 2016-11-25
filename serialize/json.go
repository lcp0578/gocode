package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Studnet struct {
	Name string
	Age  int
}

func main() {
	f, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := &Student{"lcpeng", 25}
	encoder := json.NewEncoder(f)
	encoder.Encode(s)
	// 重置指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := json.NewDecoder(f)
	var s1 Student
	decoder.Decode(&s1)
	fmt.Println(s1)
}
