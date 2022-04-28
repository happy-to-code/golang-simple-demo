package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// Transaction 交易结构
type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	TXid  []byte // 引用的交易ID
	Index int64  // 引用的output索引值
	sig   string // 解锁脚本
}

type TXOutput struct {
	value      float64 // 转账金额
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
