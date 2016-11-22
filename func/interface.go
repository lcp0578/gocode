package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type bindFunc func(int, int) int

var jsonDoc = []byte(`["add", "sub", "mul", "div"]`)

var registry = map[string]bindFunc{
	"add": func(x, y int) int { return x + y },
	"sub": func(x, y int) int { return x - y },
	"mul": func(x, y int) int { return x * y },
	"div": func(x, y int) int { return x / y },
}

func (fn *bindFunc) UnmarshalJSON(b []byte) error {
	var name string
	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}
	found := registry[name]
	// get the function out of our function registry
	if found == nil {
		return fmt.Errorf("Unknow function in (*bindFunc) UnmarshalJSON: %s", name)
	}
	// dereference the pointer receiver, so that the change are visible to the caller
	*fn = found
	return nil
}

func main() {
	rand.Seed(time.Now().Unix())
	var fns []bindFunc
	if err := json.Unmarshal(jsonDoc, &fns); err != nil {
		log.Fatal(err)
	}
	fn := fns[rand.Intn(len(fns))]
	x := fn(12, 5)
	fmt.Println(x)
}
