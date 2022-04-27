package main

import "fmt"

func (cli *CLI) PrintBlockChain() {
	bc := cli.bc
	// 创建迭代器
	it := bc.NewIterator()
	for true {
		block := it.Next()
		fmt.Printf("版本号：%d\n", block.Version)
		fmt.Printf("prev区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
		fmt.Printf("时间戳：%d\n", block.TimeStamp)
		fmt.Printf("难度值：%d\n", block.Difficulty)
		fmt.Printf("时随机数：%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Data)
		fmt.Println("==================================")
		fmt.Println()

		if len(block.PrevHash) == 0 {
			fmt.Println("------------遍历结束------------")
			break
		}
	}
}

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("添加区块成功")
}
