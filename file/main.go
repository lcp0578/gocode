package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(path.Base("/user/local"))
	fmt.Println(path.Base("D:\\windows")) // 无法识别windows路径分隔符
	fmt.Println(path.Base(strings.Replace("D:\\windows", "\\", "/", -1)))
	fmt.Println(path.Clean("/user/../../root"))
	fmt.Println(path.Dir("/root/lcp0578/a"))
	fmt.Println(path.Dir("/root/lcp0578/../lcpeng/b"))
	fmt.Println(path.Ext("/root/test"))
	fmt.Println(path.Ext("/root/test.php"))
	fmt.Println(path.IsAbs("/root/test"))
	fmt.Println(path.Join("/root", "lcpeng", "lcp"))
	fmt.Println(path.Split("/root/lcpeng/main.go"))

	fmt.Println(filepath.Abs("./test/main.go"))
}
