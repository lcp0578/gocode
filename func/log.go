package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	do := func(n int) bool {
		path := fmt.Sprintf("../log/%d.txt", n)
		f, err := os.Open(path)

		if err != nil {
			log.Println(err)
			fmt.Println(err.Error())
			return false
		}

		// // 延迟调用，释放资源
		defer f.Close()

		// // do something
		buf := make([]byte, 1024)
		var content string
		for {
			n, err := f.Read(buf)
			if err != nil {
				log.Println(err)
				continue
			}
			if n == 0 {
				break
			}
			content += string(buf[0:n])
		}
		fmt.Println(content)
		return true
	}

	for i := 0; i < 10; i++ {
		result := do(i)
		fmt.Println(result)
	}
}
