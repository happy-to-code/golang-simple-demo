package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const (
	blockChainDB  = "blockChain.db"
	blockBucket   = "blockBucket"
	lastBlockHash = "lastBlockHash"
)

type BlockChain struct {
	// blocks []*Block
	db   *bolt.DB
	tail []byte // 最后一个块的hash值
}

func NewBlockChain() *BlockChain {
	var lastHash []byte
	// 	1、打开数据库
	db, err := bolt.Open("C:\\Users\\yida\\GolandProjects\\GoProjectDemo\\btc-demo\\"+blockChainDB, 0600, nil)
	// defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(blockBucket))
		if err != nil {
			log.Panic("创建bucket失败")
		}

		// 获取创世区块
		genesisBlock := GenesisBlock()
		// 写数据
		bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
		bucket.Put([]byte(lastBlockHash), genesisBlock.Hash)

		lastHash = bucket.Get([]byte(lastBlockHash))
		return nil
	})

	return &BlockChain{
		db:   db,
		tail: lastHash,
	}
}
func GenesisBlock() *Block {
	return NewBlock("我是创世区块", []byte{})
}

func (bc *BlockChain) AddBlock(data string) {
	db := bc.db
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket不可以为空，请检查")
		}

		block := NewBlock(data, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(lastBlockHash), block.Hash)

		// 更新内存中的区块链tail
		bc.tail = block.Hash

		return nil
	})
}
