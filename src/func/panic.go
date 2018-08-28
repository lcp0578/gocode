package main

import (
	"fmt"
	"time"
)

/**
    panic (主动爆出异常) 与 recover （收集异常）
    必须注意:
    1.defer 需要放在 panic 之前定义，另外recover只有在 defer 调用的函数中才有效。
    2.recover处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点.
    3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用

    recover 用来对panic的异常进行捕获. panic 用于向上传递异常，执行顺序是在 defer 之后。

**/
func main() {
	fmt.Println("main start")
	test()
	fmt.Println("main end")
}

func test() {
	fmt.Println("test func start")
	defer func() {
		fmt.Println("test defer start")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//panicFunc()
	for {
		fmt.Println("test for start")
		a := []string{"a", "b"}
		fmt.Println(a[3])
		panic("test for panic")
		time.Sleep(1 * time.Second)
		fmt.Println("test for end")
	}
	fmt.Println("test func start")
}
func panicFunc() {
	panic("panic func")
}
