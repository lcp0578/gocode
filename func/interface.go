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

/**
摘自encoding/json/decode.go

// Unmarshaler is the interface implemented by types
// that can unmarshal a JSON description of themselves.
// The input can be assumed to be a valid encoding of
// a JSON value. UnmarshalJSON must copy the JSON data
// if it wishes to retain the data after returning.
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}
//如果你定义了这个方法，  encoding/json包就会使用它去解析json数据
**/
func (fn *bindFunc) UnmarshalJSON(b []byte) error {
	var name string
	if err := json.Unmarshal(b, &name); err != nil {
		return err
	}
	fmt.Println("name:", name)
	found := registry[name]
	fmt.Println("found", found)
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
	fmt.Println(fns)
	fn := fns[rand.Intn(len(fns))]
	fmt.Println(fn)
	x := fn(12, 5)
	fmt.Println(x)
}
