package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db                 *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db:                 bc.db,
		currentHashPointer: bc.tail,
	}
}

func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panicln("迭代器中bucket为空，请检查")
		}
		blockTem := bucket.Get(it.currentHashPointer)
		block = Deserialize(blockTem)

		// 指针左移
		it.currentHashPointer = block.PrevHash

		return nil
	})
	return &block
}
