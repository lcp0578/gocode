package main

import (
	"fmt"
	"math/rand"
	"time"
)

type binFunc func(int, int) int

type op struct {
	name string
	fn   func(int, int) int
}

func main() {
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	// create a slice of functions
	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
	}
	// pick one of those functions at random
	fn := fns[rand.Intn(len(fns))]
	fmt.Println(fn(3, 4))

	// create a slice of ops
	ops := []op{
		{"add", func(x, y int) int { return x + y }},
		{"sub", func(x, y int) int { return x - y }},
		{"mul", func(x, y int) int { return x * y }},
		{"div", func(x, y int) int { return x / y }},
		{"mod", func(x, y int) int { return x % y }},
	}
	o := ops[rand.Intn(len(ops))]
	x, y := 6, 3
	fmt.Println(o.name, x, y)
	fmt.Println(o.fn(x, y))
}
