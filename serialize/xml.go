package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int
}

//118896ead6dcc949fb8982dcd0dcb9c0ce54e05a
//94f400ef1410b9d04d039e54467870a8c42f24fb
//637f84fb9546c43931e2f93f3b3cfba9118dde37
func main() {
	f, err := os.Create("date.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	s := &Student{"lcpeng", 25}
	// 创建encoder对象
	encoder := xml.NewEncoder(f)
	encoder.Encode(s)
	// 重置指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := xml.NewDecoder(f)
	// 可以把一个对像直接序列化到io.Writer
	var s1 Student
	// io.Reader 中，返序列化xml
	decoder.Decode(&s1)
	fmt.Println(s1)

	// 直接序列化
	result, err := xml.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))
}
