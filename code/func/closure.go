package main

func main() {
	func(s string) {
		println(s)
	}("hello lcpeng")

	add := func(x, y int) int {
		return x + y
	}
	println(add(1, 2))

	add2 := test()
	println(add2(1, 2))

	testStruct()
	testChannel()

	//func() {} // 匿名函数必须调用，func literal evaluated but not used

	f := testClosure(0123)
	f()
	fa, fb := testClosure2(0x100)
	//fa()
	fb()
	fa()

	for _, fa := range testFuncArr() {
		fa()
	}
}

func test() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}
	x := calc{
		mul: func(x, y int) int {
			return x + y
		},
	}
	println(x.mul(2, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(x, y int) int {
		return x + y
	}
	println((<-c)(1, 2))
}

func testClosure(x int) func() {
	println(&x)
	return func() {
		println(&x, x)
	}
}

func testClosure2(x int) (func(), func()) {
	println(&x)
	return func() {
			println(&x, x)
		}, func() {
			x += 10
			println(&x, x)
		}
}

func testFuncArr() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			//println(&i, i)
			println(&x, x)
		})
	}
	return s
}
