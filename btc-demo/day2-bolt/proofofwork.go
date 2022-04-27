package main

import (
	"bytes"
	"crypto/sha256"
	"log"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	// 临时设置一个难度值
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	temInt := big.Int{}
	temInt.SetString(targetStr, 16)

	pow.target = &temInt

	return &pow
}

func (pow *ProofOfWork) Run() ([]byte, uint64) {
	block := pow.block
	var nonce uint64
	var hash [32]byte
	for {
		// 1、拼装数据   区块数据+变化的随机数
		temp := [][]byte{
			Uint642Byte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint642Byte(block.TimeStamp),
			Uint642Byte(block.Difficulty),
			Uint642Byte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(temp, []byte{})
		// 2、做hash运算
		hash = sha256.Sum256(blockInfo)
		// 3、与pow中的target进行比较
		// 3.1将hash转换下
		temInt := big.Int{}
		temInt.SetBytes(hash[:])
		// 	比较
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if temInt.Cmp(pow.target) == -1 { // 找到了
			log.Printf("挖矿成功！hash:%x,nonce:%d\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}
}
