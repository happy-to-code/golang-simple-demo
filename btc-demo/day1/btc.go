package main

import (
	"crypto/sha256"
	"fmt"
)

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

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	// 获取创世区块
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}
func GenesisBlock() *Block {
	return NewBlock("我是创世区块", []byte{})
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	preHash := lastBlock.Hash

	block := NewBlock(data, preHash)
	bc.blocks = append(bc.blocks, block)
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("高度为1的区块")
	bc.AddBlock("高度为2的区块")
	for i, block := range bc.blocks {
		fmt.Printf("===============================当前区块高度：%d=============================\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PreHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Data)
	}
}
