package main

import (
	"fmt"
	"sync"
)

// err 错误示例
func main() {
	var count = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
}
