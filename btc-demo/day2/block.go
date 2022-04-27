package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
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
	Data       []byte
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

func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   preBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	// 设置hash值
	// block.SetHash()

	// 创建pow对象
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}
func (block *Block) SetHash() {
	// 方法一拼装数据
	// var blockInfo []byte
	// blockInfo = append(blockInfo, Uint642Byte(block.Version)...)
	// blockInfo = append(blockInfo, block.PrevHash...)
	// blockInfo = append(blockInfo, block.MerkelRoot...)
	// blockInfo = append(blockInfo, Uint642Byte(block.TimeStamp)...)
	// blockInfo = append(blockInfo, Uint642Byte(block.Difficulty)...)
	// blockInfo = append(blockInfo, Uint642Byte(block.Nonce)...)
	// blockInfo = append(blockInfo, block.Data...)

	// 方法二拼装数据
	temp := [][]byte{
		Uint642Byte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint642Byte(block.TimeStamp),
		Uint642Byte(block.Difficulty),
		Uint642Byte(block.Nonce),
		block.Data,
	}
	blockInfo := bytes.Join(temp, []byte{})

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
