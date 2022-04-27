package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("高度为1的区块")
	bc.AddBlock("高度为2的区块")
	// for i, block := range bc.blocks {
	// 	fmt.Printf("%d===============================当前区块高度：%d=============================\n", i, i)
	// 	fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
	// 	fmt.Printf("当前区块哈希值：%x\n", block.Hash)
	// 	fmt.Printf("区块数据：%s\n", block.Data)
	// }

	// 创建迭代器
	it := bc.NewIterator()
	for true {
		block := it.Next()
		fmt.Printf("prrev区块哈希值：%x\n", block.PrevHash)
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
