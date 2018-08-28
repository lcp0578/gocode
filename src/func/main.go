package main

func main() {
	f := hello
	exec(f)
}

func hello() {
	println("hello golang")
	println(a == nil)
	//println(a == b)  //invalid operation: a == b (func can only be compared to nil)
	var ta *int = test()
	println(*ta, ta)
}

func exec(f func()) {
	f()
}

func a() {}
func b() {}

func test() *int {
	a := 0x100
	return &a
}
