package main

import (
	"fmt"
)

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	if a == 0 && b == 2 {
		err := fmt.Errorf("%s", "----")

		panic(err)
	}
	print(b, " ")
	print(a, "\n")
}

func main() {
	for true {
		go f()
		g()
	}
}

// 0 1
// 2 1
// 0 0
