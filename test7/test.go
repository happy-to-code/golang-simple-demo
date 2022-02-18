package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strList []string
	for i := 1; i <= 160; i++ {
		strList = append(strList, strconv.FormatInt(int64(i), 10))
	}
	fmt.Println(strList)
	fmt.Println("==========================================================")

	list := cutOffList(strList, 80)
	fmt.Println(list)

}

//strList 要切分的数组
//listSize 切分后每个数组的size
func cutOffList(strList []string, listSize int) (ss [][]string) {
	// 对listSize取模
	mod := len(strList) % listSize
	// 对listSize取余
	k := len(strList) / listSize

	// 计算循环的截止数
	var end int
	if mod == 0 {
		end = k
	} else {
		end = k + 1
	}

	for i := 0; i < end; i++ {
		if i != k {
			ss = append(ss, strList[i*listSize:(i+1)*listSize])
		} else {
			ss = append(ss, strList[i*listSize:])
		}
	}
	return
}
