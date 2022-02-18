package main

import "fmt"

func main() {
	str := "chinese"
	city := "beijing"
	zstr := []byte(str)
	for i := 0; i < 100000; i++ {
		copy(zstr, city)
	}
	fmt.Println(string(zstr))
}
