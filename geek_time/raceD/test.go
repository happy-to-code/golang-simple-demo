package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a *int) int {
	return *a + 1
}
func main() {
	var t int = 0
	var lock sync.Mutex

	lock.Lock()
	defer lock.Unlock()
	for i := 0; i < 10000; i++ {
		go add(&t)
		t++
	}

	fmt.Println("===>", t)
	time.Sleep(time.Second * 3)
}
