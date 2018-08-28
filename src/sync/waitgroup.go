package main

import (
	"fmt"
	"net/http"
	"sync"
)

/**
 * sync.WaitGroup只有3个方法，Add()，Done()，Wait()。
 * 其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
 * A WaitGroup waits for a collection of goroutines to finish.
 * The main goroutine calls Add to set the number of goroutines to wait for.
 * Then each of the goroutines runs and calls Done when finished.
 * At the same time, Wait can be used to block until all goroutines have finished.
 **/
func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.baidu.com",
		"http://kitcloud.cn",
		"http://www.golang.org",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
			fmt.Println(url)
		}(url)
	}
	wg.Wait()
	fmt.Println("over")
}
