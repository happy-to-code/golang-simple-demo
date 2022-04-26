package main

import "fmt"

func main() {

	ids := []int{1, 3, 56, 8}
	var s string
	for i, id := range ids {
		if i == 0 {
			s += fmt.Sprintf("%d", id)
		} else {
			s += fmt.Sprintf("%s%d", ",", id)
		}
	}
	fmt.Println(s)

}
