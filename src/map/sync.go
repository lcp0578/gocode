package main

import (
	_ "fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.RWMutex // 使用读写锁，获得最佳性能
	m := make(map[string]int)

	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			// fmt.Println(m)
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			lock.RLock()
			_ = m["b"]
			// fmt.Println(m)
			lock.RUnlock()
			time.Sleep(time.Microsecond)
		}
	}()
	select {}
}
