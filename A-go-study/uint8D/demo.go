package main

import "fmt"

func main() {
	var i int8 = 127
	fmt.Println(i)
	fmt.Println(i + 1)
	fmt.Println(i * i)

	fmt.Println("---------------------")
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 255 0 1

	fmt.Println("--------------------")
	fmt.Println(3 << 1)
	fmt.Println(3 << 2)
	fmt.Println(3 << 3)

}
