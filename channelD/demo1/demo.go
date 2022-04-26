package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func printer(s string) {
	lock.Lock()
	defer lock.Unlock()
	for _, i := range s {
		fmt.Printf("%c", i)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
}
func person(s string) {
	printer(s)
}

func main() {
	go person("hello")
	go person("world")

	for true {

	}

}
