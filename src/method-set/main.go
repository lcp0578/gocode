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

type Car interface {
	Run()
}

func Driver(car Car) {
	car.Run()
}

type Audi struct {
	Name  string
	Price float32
	Color string
}

func (audi Audi) Run() {
	fmt.Println("Audi:", audi.Name, audi.Price, "w", audi.Color, " run")
}

type Benz struct {
	Name  string
	Price float32
	Color string
}

func (benz *Benz) Run() {
	fmt.Println("Benz:", benz.Name, benz.Price, "w", benz.Color, " run")
}

type GlkBenz struct {
	Benz
	num string
}

type QAudi struct {
	Audi
	num string
}

type BlackCat struct {
	Cat
	Age int
}

func main() {
	cat1 := Cat{"cat1", "black"}
	cat1.Meow()
	cat1.Miao()
	cat2 := &Cat{"cat2", "white"}
	cat2.Meow()
	cat2.Miao()
	//结论：普通情况下，类型 T 和 *T 上的方法集是互相继承的。

	a6 := Audi{"A6", 30.60, "black"}
	Driver(a6)
	a8 := &Audi{"A8", 62.56, "white"}
	Driver(a8)

	//gla200 := Benz{"gla200", 31.20, "black"}
	//Driver(gla200)
	/**
	  cannot use gla200 (type Benz) as type Car in argument to Driver:
	  Benz does not implement Car (Run method has pointer receiver)
	  **/
	s300 := &Benz{"s300", 200.00, "black"}
	Driver(s300)
	/**
	   在接口中的method，对于普通类型T：
	   T的methods set里不会继承包含*T实现的method，除非T自己实现相对应的method。
	   但是，*T会继承T的method set。
	**/
	bc1 := BlackCat{
		Cat{"bc1", "black"},
		3,
	}
	bc2 := &BlackCat{
		Cat: Cat{Name: "bc2", Color: "black"},
		Age: 5}
	bc1.Meow()
	bc1.Miao()
	bc1.Cat.Meow()
	bc1.Cat.Miao()
	bc2.Meow()
	bc2.Miao()
	bc2.Cat.Meow()
	bc2.Cat.Miao()
	/**
	  嵌入类型的类型中，外部类型自己未曾实现的methods被携带的内部函数实现时，
	  外部类型也会将这些methods加入到自己的methods set里。
	  **/
	glk := &GlkBenz{
		Benz{"glk", 405.00, "black"},
		"glk300",
	}
	Driver(glk)
	q7 := QAudi{
		Audi{"Q7", 190, "white"},
		"q7",
	}
	Driver(q7)
	q8 := &QAudi{
		Audi{"Q8", 230, "white"},
		"q8",
	}
	Driver(q8)
}
