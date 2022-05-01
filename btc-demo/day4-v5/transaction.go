package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	// "github.com/btcsuite/btcutil/base58"
	"GoProjectDemo/btc-demo/day4-v5/lib/base58"
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
	// Sig   string // 解锁脚本
	Signature []byte // 真正的数字签名  由r s 拼成的[]byte
	PubKey    []byte
}

type TXOutput struct {
	Value float64 // 转账金额
	// PubKeyHash string  // 锁定脚本
	PubKeyHash []byte // 收款方的公钥哈希
}

func (output *TXOutput) Lock(address string) {
	// 解码
	// 	截取出公钥哈希：去除version（1字节），去除校验码（4字节）
	addressByte := base58.Decode(address)
	len := len(addressByte)
	pubKeyHash := addressByte[1 : len-4]

	output.PubKeyHash = pubKeyHash
}

func NewTXOutput(value float64, address string) *TXOutput {
	output := TXOutput{
		Value: value,
	}
	// 锁定地址
	output.Lock(address)

	return &output
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
		TXid:      []byte{},
		Index:     -1,
		Signature: nil,
		PubKey:    []byte(data),
	}
	// output := TXOutput{
	// 	Value:      reward,
	// 	PubKeyHash: address,
	// }
	output := NewTXOutput(reward, address)

	// 对于挖矿交易来说  只有一个input 和 一个 output
	tx := Transaction{
		TXID:      []byte{},
		TXInputs:  []TXInput{input},
		TXOutputs: []TXOutput{*output},
	}

	// SetHash 设置交易ID
	tx.SetHash()

	return &tx
}

// NewTransaction
// 创建普通转账交易
func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	ws := NewWallets()
	wallet := ws.WalletsMap[from]
	if wallet == nil {
		return nil
	}
	pubKey := wallet.PubKey
	// private := wallet.Private

	pubKeyHash := HashPubKey(pubKey)
	// 1.找到合理的UTXO集合  map[string][]int64
	utxos, resValue := bc.FindNeedUTXOs(pubKeyHash, amount)
	if resValue < amount {
		log.Panicf("【%f】【%f】余额不足，交易失败\n", resValue, amount)
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput
	// 2.创建交易输入  将这些UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), nil, pubKey}
			inputs = append(inputs, input)
		}
	}

	// 3.创建交易输出  创建outputs
	output := NewTXOutput(amount, to)
	outputs = append(outputs, *output)

	// 4.如果有找零  找零
	if resValue > amount {
		// 找零
		outputs = append(outputs, *NewTXOutput(resValue-amount, from))
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
