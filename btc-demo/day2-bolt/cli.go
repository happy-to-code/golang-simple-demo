package main

import (
	"fmt"
	"os"
)

//  这是一个用来接收命令行参数并且控制区块去了操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA	"add data to blockChain"
	printChain				"print all blockChain data"
`

func (cli *CLI) Run() {
	// 	获取命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		fmt.Println()
		return
	}

	// 	分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块使用参数不当，请检查")
			fmt.Printf(Usage)
		}
	case "printChain":
		fmt.Println("打印区块信息")
		cli.PrintBlockChain()
	default:
		fmt.Println("无效的参数")
		fmt.Printf(Usage)
	}
}
