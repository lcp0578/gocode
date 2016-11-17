package main

import (
	"fmt"
	"os"
)

func main() {
	f, openErr := os.OpenFile("D:\\gowork\\src", os.O_RDONLY, 0666)
	if openErr != nil {
		fmt.Println(openErr.Error())
		return
	}
	arrFile, readErr := f.Readdir(0)
	if readErr != nil {
		fmt.Println(readErr.Error())
		return
	}
	for k, v := range arrFile {
		fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
	}
	// GVfmumMP_ksetegQQIOlZIVVIAFbhSR4xHT5kGAGsoQ
}
