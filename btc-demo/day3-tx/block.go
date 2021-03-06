package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Version    uint64
	PrevHash   []byte
	MerkelRoot []byte
	TimeStamp  uint64
	Difficulty uint64
	Nonce      uint64
	Hash       []byte
	// Data       []byte
	Transactions []*Transaction
}

// Uint642Byte 将uint64 ==>[]byte
func Uint642Byte(i uint64) (by []byte) {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, i)
	if err != nil {
		log.Panicln(err)
	}
	by = buffer.Bytes()
	return
}

func NewBlock(txs []*Transaction, preBlockHash []byte) *Block {
	block := Block{
		Version:      00,
		PrevHash:     preBlockHash,
		MerkelRoot:   []byte{},
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulty:   0,
		Nonce:        0,
		Hash:         []byte{},
		Transactions: txs,
	}
	// 设置梅克尔根
	block.MerkelRoot = block.MakeMerkelRoot()

	// 设置hash值
	// block.SetHash()

	// 创建pow对象
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panicf("[%v]编码出错,err:%s\n", block, err.Error())
	}
	return buffer.Bytes()
}

func Deserialize(data []byte) Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Printf("[%s]解码失败,err:%s", data, err.Error())
	}
	return block
}

// MakeMerkelRoot 模拟梅克尔根
func (block *Block) MakeMerkelRoot() []byte {
	var info []byte
	for _, tx := range block.Transactions {
		info = append(info, tx.TXID...)
	}
	hash := sha256.Sum256(info)
	return hash[:]
}
