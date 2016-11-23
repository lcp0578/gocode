package main

import (
	"fmt"
	"math/rand"
	"time"
)

var delay = 200 * time.Millisecond

func pickFunc(fns ...anonIntFunc) anonIntFunc {
	return fns[rand.Intn(len(fns))]
}

func produce(c chan anonIntFunc, n int, fns ...anonIntFunc) {
	defer close(c)
	for i := 0; i < n; i++ {
		c <- pickFunc(fns...)
	}
}

type anonIntFunc func(int) int

func main() {
	rand.Seed(time.Now().Unix())

	x := 10
	fns := []anonIntFunc{
		func(x int) int {
			x += 1
			return x
		},
		func(x int) int {
			x += 2
			return x
		},
		func(x int) int {
			x += 3
			return x
		},
		func(x int) int {
			x += 4
			return x
		},
		func(x int) int {
			x += 5
			return x
		},
		func(x int) int {
			x += 6
			return x
		},
		func(x int) int {
			x += 7
			return x
		},
		func(x int) int {
			x += 8
			return x
		},
		func(x int) int {
			x += 9
			return x
		},
	}
	c := make(chan anonIntFunc)

	go produce(c, x, fns...)

	for fn := range c {
		x = fn(x)
		fmt.Println(x)
		time.Sleep(delay)
	}
}
