package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "(设备1,TBAeMeN92z8FjoB4UuAFU1cgmN95TFcyVo)"
	if strings.HasPrefix(s, "(") {
		s = s[1:]
	}
	if strings.HasSuffix(s, ")") {
		s = s[:strings.LastIndex(s, ")")]
	}
	fmt.Println(s)

	s2 := "(设备1,TBAeMeN92z8FjoB4UuAFU1cgmN95TFcyVo)"
	s2 = s2[strings.Index(s2, "(")+1 : strings.LastIndex(s2, ")")]
	fmt.Println(s2)

}
