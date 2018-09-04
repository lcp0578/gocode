package main

import (
	"errors"
	"fmt"
	"sync"
)

type MyMap struct {
	mp    map[string]int
	mutex *sync.RWMutex
}

func (this *MyMap) Get(key string) (int, error) {
	this.mutex.Lock()
	i, ok := this.mp[key]
	this.mutex.Unlock()
	if !ok {
		return i, errors.New("key not exist")
	}
	return i, nil
}

func (this *MyMap) Set(key string, val int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.mp[key] = val
}

func (this *MyMap) Display() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	for k, v := range this.mp {
		fmt.Println(k, "=>", v)
	}
}
