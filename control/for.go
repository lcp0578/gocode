package main

func main() {
	for i, c := 0, count(); i < c; i++ { // 初始化语句，count仅执行一次
		println("a", i)
	}

	c := 0
	for c < count() { // 重复执行
		println("b", c)
		c++
	}
}

func count() int {
	print("count func")
	return 3
}
