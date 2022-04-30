package main

import (
	"fmt"
	"os"
	"strconv"
)

//  这是一个用来接收命令行参数并且控制区块去了操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	printChain				"print all blockChain data"
	getBalance --address ADDRESS	"获取地址余额"
	send FROM TO AMOUNT MINER DATA 	"由FROM转AMOUNT给TO，由MINER挖矿2，同时写入DATA"
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
	case "printChain":
		fmt.Println("打印区块信息")
		cli.PrintBlockChain()
	case "getBalance":
		fmt.Println("获取余额")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cli.GetBalance(address)
		}
	case "send":
		fmt.Println("转账开始…………")
		if len(args) != 7 {
			fmt.Println("参数错误，请检查")
			fmt.Printf(Usage)
			fmt.Println()
			return
		}
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64)
		miner := args[5]
		data := args[6]
		cli.Send(from, to, amount, miner, data)
		fmt.Println("转账结束…………")

	default:
		fmt.Println("无效的参数")
		fmt.Printf(Usage)
	}
}
