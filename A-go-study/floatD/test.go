package main

import (
	"fmt"
	"math"
)

func main() {
	// 一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精度；
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(f == f+1) // true
	fmt.Println("--------")
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
