package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const reward = 50

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

// NewTransaction
// 创建普通转账交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 1.找到合理的UTXO集合  map[string][]int64
	utxos, resValue := bc.FindNeedUTXOs(from, amount)
	if resValue < amount {
		log.Panicf("【%f】【%f】余额不足，交易失败\n", resValue, amount)
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput
	// 2.创建交易输入  将这些UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			inputs = append(inputs, TXInput{
				TXid:  []byte(id),
				Index: i,
				Sig:   from,
			})
		}
	}

	// 3.创建交易输出  创建outputs
	outputs = append(outputs, TXOutput{
		Value:      amount,
		PubKeyHash: to,
	})

	// 4.如果有找零  找零
	if resValue > amount {
		// 找零
		outputs = append(outputs, TXOutput{
			Value:      resValue - amount,
			PubKeyHash: from,
		})
	}

	tx := Transaction{
		TXID:      []byte{},
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
	tx.SetHash()
	return &tx
}

// IsCoinbase 判断当前交易是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	// 交易的input只有1个  && 交易ID为空   && 交易的index为-1
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}
