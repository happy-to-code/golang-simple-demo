package main

import (
	"fmt"
)

func main() {
	c := make(chan struct{})
	go func() {
		fmt.Println("111")
		c <- struct{}{}
	}()
	<-c
}
