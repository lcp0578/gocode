package main

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}

func testCap() map[int]int {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCap()
	}
}
