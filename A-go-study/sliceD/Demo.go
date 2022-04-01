package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s, 2)) // "[5 6 8 9]"
}
func remove(slice []int, i int) []int {
	fmt.Println("1-->", slice[i:])
	fmt.Println("2-->", slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Println("3-->", slice)

	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) (a []int) {
	a = append(slice[:i], slice[i+1:]...)
	return
}
