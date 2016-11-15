package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Color string
}

func (c Cat) Meow() {
	fmt.Println(c.Name, c.Color)
}

func (c *Cat) Miao() {
	fmt.Println("Miao~", c.Name, " Miao~", c.Color)
}

//普通情况下，类型 T 和 *T 上的方法集是互相继承的。
func main() {
	cat1 := Cat{"cat1", "black"}
	cat1.Meow()
	cat1.Miao()
	cat2 := &Cat{"cat2", "white"}
	cat2.Meow()
	cat2.Miao()
}
