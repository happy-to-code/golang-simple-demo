package main

import (
	"fmt"
	"strings"
)

func main() {
	smap := make(map[string]string)

	var s = "  tag, tag3,tag3，tag4"
	s = strings.TrimSpace(s)
	if strings.Contains(s, "，") {
		s = strings.ReplaceAll(s, "，", ",")
	}

	if strings.Contains(s, ",") {
		split := strings.Split(s, ",")
		for _, s2 := range split {
			s2 = strings.TrimSpace(s2)
			if len(s2) > 0 {
				smap[s2] = s2
			}
		}
	} else {
		if len(s) > 0 {
			smap[s] = s
		}
	}

	// fmt.Println("-----------------------------------------------")
	for k, v := range smap {
		fmt.Printf("==========>%s<==============%s\n", k, v)
	}
}
