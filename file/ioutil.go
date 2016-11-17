package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	arrFile, err := ioutil.ReadDir("D:\\gowork\\src\\gopkg.in\\yaml.v2")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for k, v := range arrFile {
		fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
	}
}
