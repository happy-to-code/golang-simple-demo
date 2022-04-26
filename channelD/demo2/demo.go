package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(s string) {

	for _, i := range s {
		fmt.Printf("%c", i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}
func person1(s string) {
	printer(s)
	ch <- 561
}
func person2(s string) {
	<-ch
	printer(s)
}
func main() {
	go person1("hello")
	go person2("world")

	for true {

	}

}
