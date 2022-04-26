package main

import "crypto/sha256"

type Block struct {
	PreHash []byte
	Hash    []byte
	Data    []byte
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		PreHash: preBlockHash,
		Hash:    []byte{},
		Data:    []byte(data),
	}
	// 设置hash值
	block.SetHash()
	return &block
}
func (block *Block) SetHash() {
	blockInfo := append(block.PreHash, block.Data...)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
