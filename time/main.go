package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) //时间原点
	d, err := time.Parse("2006-01-02 15:04:05", "2016-11-11 22:46:20")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)
}
