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
	db   *bolt.DB
	tail []byte // 最后一个块的hash值
}

func NewBlockChain(address string) *BlockChain {
	var lastHash []byte
	// 	1、打开数据库
	db, err := bolt.Open(blockChainDB, 0600, nil)
	// defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}
	db.Update(func(tx *bolt.Tx) error {
		// 找到抽屉bucket(如果没有，就创建）
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			// 没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建bucket失败")
			}

			// 创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock(address)

			// 3. 写数据
			// hash作为key， block的字节流作为value，尚未实现
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte(lastBlockHash), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte(lastBlockHash))
		}

		return nil
	})

	return &BlockChain{
		db:   db,
		tail: lastHash,
	}
}
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTx(address, "---我是创世区块---")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	db := bc.db
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket不可以为空，请检查")
		}

		block := NewBlock(txs, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(lastBlockHash), block.Hash)

		// 更新内存中的区块链tail
		bc.tail = block.Hash

		return nil
	})
}

func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var utxo []TXOutput

	return utxo
}
