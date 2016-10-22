package main

import (
	"fmt"
)

type FruitName func() string

type Fruit struct {
	getFruitName FruitName
}

type Apple struct {
	Fruit
}

func (this Fruit) DisplayName() {
	fmt.Println(this.getName())
}

func (this Fruit) getName() string {
	return "水果"
}
func (this Apple) DisplayName() {
	fmt.Println(this.getName())
}
func (this Apple) getName() string {
	return "苹果"
}

func NewFruit() *Fruit {
	f := new(Fruit)
	return f
}

func NewApple() *Apple {
	a := new(Apple)
	return a
}
func main() {
	f := NewFruit()
	f.DisplayName()
	a := NewApple()
	a.DisplayName()
}
