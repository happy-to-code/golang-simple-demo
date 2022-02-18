package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("hello world")
	fmt.Println("hello world3")
	handledBlockNum := selectMaxHasHandledBlockNum()

	for i := handledBlockNum; ; i++ {
		fmt.Printf("-->%d--%T\n", i, i)

		if i == 2453131 {
			//panic("ddfdfd")
			return
		}
		fmt.Println("===============")

	}
	fmt.Println("pppp")
}

func selectMaxHasHandledBlockNum() int64 {
	return 2453129
}
