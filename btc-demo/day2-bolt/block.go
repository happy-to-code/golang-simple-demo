package main

import (
	"bytes"
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

func (block *Block) toByte() []byte {

	// todo
	return []byte{}
}
