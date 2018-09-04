package main

func main() {
	x, y := 1, 2
	defer func(a int) {
		println("defer x, y =", a, y)
	}(x)
	x += 100
	y += 200
	println(x, y)

	//多个延迟调用按FILO次序执行
	defer println("a")
	defer println("b")
	defer func() {
		println("c")
		// fmt.Println("c")
	}()
	t := test(0xf00)
	println(t)
}

func test(z int) int {
	defer func() {
		println("defer1", z)
		z += 100
		println("defer2", z)
	}()
	return z
}
