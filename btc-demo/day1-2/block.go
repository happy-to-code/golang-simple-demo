package main

import (
	"crypto/sha256"
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
	block.SetHash()
	return &block
}
func (block *Block) SetHash() {
	var blockInfo []byte
	blockInfo = append(blockInfo, Uint642Byte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint642Byte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint642Byte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint642Byte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Data...)

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
