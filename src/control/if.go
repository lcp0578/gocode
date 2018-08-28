package main

import (
	"errors"
	"log"
)

func check(x int) error {
	if x <= 0 {
		return errors.New("x <= 0")
	}
	return nil
}

func main() {
	x := 0
	if err := check(x); err == nil {
		x++
		println(x)
	} else {
		log.Fatalln(err)
	}
}
