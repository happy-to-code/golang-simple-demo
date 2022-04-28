package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const reward = 12.5

// Transaction 交易结构
type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	TXid  []byte // 引用的交易ID
	Index int64  // 引用的output索引值
	Sig   string // 解锁脚本
}

type TXOutput struct {
	Value      float64 // 转账金额
	PubKeyHash string  // 锁定脚本
}

// SetHash 设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panicln(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

// NewCoinbaseTx 创建挖矿交易
func NewCoinbaseTx(address string, data string) *Transaction {
	// 	挖矿交易的特点
	// 	1、只有一个in
	// 	2、无需引用交易ID
	// 	3、无需引用index
	input := TXInput{
		TXid:  []byte{},
		Index: -1,
		Sig:   data,
	}
	output := TXOutput{
		Value:      reward,
		PubKeyHash: address,
	}

	// 对于挖矿交易来说  只有一个input 和 一个 output
	tx := Transaction{
		TXID:      []byte{},
		TXInputs:  []TXInput{input},
		TXOutputs: []TXOutput{output},
	}

	// SetHash 设置交易ID
	tx.SetHash()

	return &tx
}
