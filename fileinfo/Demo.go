package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//00丨发刊词｜你的“前途”里，有你吗？
//01丨交易：货币切分了买卖，商人连接了交易

func main() {
	files, _ := ioutil.ReadDir("C:\\Users\\yida\\GolandProjects\\GoProjectDemo\\097-《刘润·商业通识30讲》")
	for _, file := range files {
		name := file.Name()
		name = name[0:strings.LastIndex(name, ".")]
		fmt.Println(name)
		//fmt.Println("================================================")
		fmt.Println("file.Name:", file.Name())
		fmt.Println("file.IsDir:", file.IsDir())
		fmt.Println("file.ModTime:", file.ModTime())
		fmt.Println("file.Mode:", file.Mode())
		fmt.Println("file.Size:", file.Size())

	}
}
