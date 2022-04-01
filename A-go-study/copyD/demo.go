package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3}
	var s2 = []int{4, 5, 6}
	copy(s2, s1)
	fmt.Println(s2)
}
