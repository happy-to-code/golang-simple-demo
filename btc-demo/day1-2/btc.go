package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("高度为1的区块")
	bc.AddBlock("高度为2的区块")
	for i, block := range bc.blocks {
		fmt.Printf("%d===============================当前区块高度：%d=============================\n", i, i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Data)
	}
}
