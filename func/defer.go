package main

import (
	"fmt"
	"os"
)

func main() {
	cont, err := ReadFile("func3.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cont)
}

func ReadFile(strFileName string) (string, error) {
	f, err := os.Open(strFileName)
	defer log("open error")
	if err != nil {
		return "", err
	}
	defer log("before")
	defer f.Close()
	defer log("after")
	buf := make([]byte, 1024)
	var strContent string = ""
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		strContent += string(buf[0:n])
	}
	return strContent, nil
}
func log(str string) {
	fmt.Println(str)
}
