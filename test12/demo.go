package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "g"
	fmt.Println("--->", t(str))
	s := "i love i is family dad ma"
	fields := strings.Fields(s)
	fmt.Println(fields)

}

func t(s string) string {
	if s == "" {
		fmt.Println("1111111111111")
		return "null"
	}
	fmt.Println("22222222222222222")
	return "not null"

}
