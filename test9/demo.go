package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		//inner:
		for j := 0; j < 4; j++ {
			fmt.Println("j:", j)
			if j == 2 {
				break
			}
		}

		fmt.Println("-->iii:", i)
	}
}
