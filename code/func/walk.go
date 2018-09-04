package main

import (
	"fmt"
	"math/rand"
	"time"
)

type walkFn func(*int) walkFn

func pickRandom(fns ...walkFn) walkFn {
	fmt.Println("pickRandom called")
	return fns[rand.Intn(len(fns))]
}

func walkEqual(i *int) walkFn {
	*i += rand.Intn(7) - 3
	fmt.Println("walkEqual called, *i=", *i)
	return pickRandom(walkForward, walkBackward)
}

func walkForward(i *int) walkFn {
	*i += rand.Intn(6)
	fmt.Println("walkForward called *i=", *i)
	return pickRandom(walkEqual, walkBackward)
}

func walkBackward(i *int) walkFn {
	*i += -rand.Intn(6)
	fmt.Println("walkBackward called *i=", *i)
	return pickRandom(walkForward, walkEqual)
}
func main() {
	rand.Seed(time.Now().Unix())
	fn, progress := walkEqual, 0
	for i := 0; i < 10; i++ {
		fn = fn(&progress)
		fmt.Println(progress)
		fmt.Println("loop:", i, " end")
	}
}
