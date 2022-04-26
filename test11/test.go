package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("--->", runtime.GOROOT())
	// fmt.Println(runtime.G)
}
