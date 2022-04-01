package main

import (
	"fmt"
	"os"
)

func main1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("=======================", os.Args)
	fmt.Println(s)
	fmt.Println("hello,世界")
	// C:\Users\yida\AppData\Local\Temp\GoLand\___go_build_GoProjectDemo_A_go_study_hello_world.exe
}
func main() {
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0])
}
