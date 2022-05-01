package main

import (
	"bytes"
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

func (bc *BlockChain) FindUTXOs(pubKeyHash []byte) []TXOutput {
	var utxo []TXOutput

	txs := bc.FindUTXOTransactions(pubKeyHash)
	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			if bytes.Equal(pubKeyHash, output.PubKeyHash) {
				utxo = append(utxo, output)
			}
		}
	}
	return utxo
}

func (bc *BlockChain) FindNeedUTXOs(senderPubKeyHash []byte, amount float64) (map[string][]int64, float64) {
	utxos := make(map[string][]int64)
	// 找到的utxo里面包含的钱总数
	var calc float64

	txs := bc.FindUTXOTransactions(senderPubKeyHash)
	for _, tx := range txs {
		for i, output := range tx.TXOutputs {
			if bytes.Equal(senderPubKeyHash, output.PubKeyHash) {
				// utxo = append(utxo, output)
				// 	找到自己最少的utxo
				// 	3.比较一下是否满足转账需求
				// 		a.满足  直接返回
				// 		b.不满足  继续统计
				if calc < amount {
					// 	1.把utxo加进来
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], int64(i))
					// 	2.统计当前utxo总额
					calc += output.Value

					// 	加完之后判断
					if calc >= amount {
						return utxos, calc
					}
				}
			}
		}
	}

	return utxos, calc
}

func (bc *BlockChain) FindUTXOTransactions(senderPubKeyHash []byte) []*Transaction {
	var txs []*Transaction // 存储所有包含utxo交易的集合
	// 定义一个map来保存消费过的output   key:output的交易ID   Value:这个交易中索引的数组
	spentOutPuts := make(map[string][]int64)

	it := bc.NewIterator()
	for {
		// 1、遍历区块
		block := it.Next()
		// 2、遍历交易
		for _, tx := range block.Transactions {
			// fmt.Printf("current txid :%x\n", tx.TXID)
			// 3、遍历output，找到和自己相关的utxo（在添加output之前检查一下是否已经消耗过了）
		OUTPUT:
			for i, output := range tx.TXOutputs {
				// fmt.Printf("current index :%d\n", i)

				// 在这里做过滤，将所有消耗过的outputs和当前的 所即将添加的output对比，如果相同 则跳过  否则添加
				int64s, ok := spentOutPuts[string(tx.TXID)]
				if ok {
					for _, j := range int64s {
						if int64(i) == j { // 当前output硬消耗过了  不要再添加了
							continue OUTPUT
						}
					}
				}

				if bytes.Equal(output.PubKeyHash, senderPubKeyHash) {
					txs = append(txs, tx) // 返回所有和我相关的utxo交易集合
				}
			}
			// 如果当前交易为挖矿交易   则不做遍历  直接跳过
			isCoinbase := tx.IsCoinbase()
			if !isCoinbase {
				// 4、遍历input，找到自己花费过的utxo集合（把自己消耗过的标识出来）
				for _, input := range tx.TXInputs {
					pubKeyHash := HashPubKey(input.PubKey)
					if bytes.Equal(pubKeyHash, senderPubKeyHash) { // 说明是目标地址address消耗过的output
						indexArray := spentOutPuts[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
						spentOutPuts[string(input.TXid)] = indexArray // 这边必须操作   不然可以用下面的语句   下面的语句==上面的3句
						// spentOutPuts[string(input.TXid)] = append(spentOutPuts[string(input.TXid)], input.Index)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			// fmt.Println("区块遍历完成退出")
			break
		}
	}

	return txs
}
