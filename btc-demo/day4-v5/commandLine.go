package main

import (
	"fmt"
	"time"
)

func (cli *CLI) PrintBlockChain() {
	bc := cli.bc
	// 创建迭代器
	it := bc.NewIterator()
	for true {
		block := it.Next()
		fmt.Printf("版本号：%d\n", block.Version)
		fmt.Printf("prev区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("梅克尔根：%x\n", block.MerkelRoot)
		fmt.Printf("时间戳：%s\n", time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:05:05"))
		fmt.Printf("难度值：%d\n", block.Difficulty)
		fmt.Printf("时随机数：%d\n", block.Nonce)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Transactions[0].TXInputs[0].PubKey)
		fmt.Println("==================================")
		fmt.Println()

		if len(block.PrevHash) == 0 {
			fmt.Println("------------遍历结束------------")
			break
		}
	}
}

func (cli *CLI) GetBalance(address string) {
	// 校验地址
	// isValidAddress := IsValidAddress(address)
	// if !isValidAddress {
	// 	fmt.Printf("无效的地址,address:%s\n", address)
	// 	return
	// }

	// 生成公钥hash
	pubKeyHash := GetPubKeyFromAddress(address)
	utxos := cli.bc.FindUTXOs(pubKeyHash)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("地址[%s]余额为：%f\n", address, total)
}
func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	// 校验地址
	// if !IsValidAddress(from) {
	// 	fmt.Printf("无效的地址,from:%s\n", from)
	// 	return
	// }
	// // 校验地址
	// if !IsValidAddress(to) {
	// 	fmt.Printf("无效的地址,to:%s\n", to)
	// 	return
	// }
	// // 校验地址
	// if !IsValidAddress(miner) {
	// 	fmt.Printf("无效的地址,miner:%s\n", miner)
	// 	return
	// }
	fmt.Printf("from:%s,to:%s,amount:%f,miner:%s,data:%s\n", from, to, amount, miner, data)
	// 	1.创建挖矿交易
	coinbaseTx := NewCoinbaseTx(miner, data)
	// 2.创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		fmt.Println("无效的交易")
		return
	}
	// 3.添加到区块
	cli.bc.AddBlock([]*Transaction{coinbaseTx, tx})
	fmt.Println("转账成功！")
}

func (cli *CLI) NewWallet() {
	ws := NewWallets()
	address := ws.CreateWallet()
	fmt.Printf("地址：%s\n", address)
}

func (cli *CLI) ListAddresses() {
	ws := NewWallets()
	addresses := ws.ListAllAddresses()
	for _, address := range addresses {
		fmt.Printf("地址：%s\n", address)
	}
}
