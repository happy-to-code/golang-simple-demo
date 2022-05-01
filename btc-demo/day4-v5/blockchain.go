package main

import (
	"bytes"
	"fmt"
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
	var txs []*Transaction // 存储所有包含utxo交易集合
	// 我们定义一个map来保存消费过的output，key是这个output的交易id，value是这个交易中索引的数组
	// map[交易id][]int64
	spentOutputs := make(map[string][]int64)

	// 创建迭代器
	it := bc.NewIterator()

	for {
		// 1.遍历区块
		block := it.Next()

		// 2. 遍历交易
		for _, tx := range block.Transactions {
			// fmt.Printf("current txid : %x\n", tx.TXID)

		OUTPUT:
			// 3. 遍历output，找到和自己相关的utxo(在添加output之前检查一下是否已经消耗过)
			//	i : 0, 1, 2, 3
			for i, output := range tx.TXOutputs {
				// fmt.Printf("current index : %d\n", i)
				// 在这里做一个过滤，将所有消耗过的outputs和当前的所即将添加output对比一下
				// 如果相同，则跳过，否则添加
				// 如果当前的交易id存在于我们已经表示的map，那么说明这个交易里面有消耗过的output

				// map[2222] = []int64{0}
				// map[3333] = []int64{0, 1}
				// 这个交易里面有我们消耗过得output，我们要定位它，然后过滤掉
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						// []int64{0, 1} , j : 0, 1
						if int64(i) == j {
							// fmt.Printf("111111")
							// 当前准备添加output已经消耗过了，不要再加了
							continue OUTPUT
						}
					}
				}

				// 这个output和我们目标的地址相同，满足条件，加到返回UTXO数组中
				// if output.PubKeyHash == address {
				if bytes.Equal(output.PubKeyHash, senderPubKeyHash) {
					// fmt.Printf("222222")
					// UTXO = append(UTXO, output)

					// !!!!!重点
					// 返回所有包含我的outx的交易的集合
					txs = append(txs, tx)

					// fmt.Printf("333333 : %f\n", UTXO[0].Value)
				} else {
					// fmt.Printf("333333")
				}
			}

			// 如果当前交易是挖矿交易的话，那么不做遍历，直接跳过

			if !tx.IsCoinbase() {
				// 4. 遍历input，找到自己花费过的utxo的集合(把自己消耗过的标示出来)
				for _, input := range tx.TXInputs {
					// 判断一下当前这个input和目标（李四）是否一致，如果相同，说明这个是李四消耗过的output,就加进来
					// if input.Sig == address {
					// if input.PubKey == senderPubKeyHash  //这是肯定不对的，要做哈希处理
					pubKeyHash := HashPubKey(input.PubKey)
					if bytes.Equal(pubKeyHash, senderPubKeyHash) {
						// spentOutputs := make(map[string][]int64)
						// indexArray := spentOutputs[string(input.TXid)]
						// indexArray = append(indexArray, input.Index)
						spentOutputs[string(input.TXid)] = append(spentOutputs[string(input.TXid)], input.Index)
						// map[2222] = []int64{0}
						// map[3333] = []int64{0, 1}
					}
				}
			} else {
				// fmt.Printf("这是coinbase，不做input遍历！")
			}
		}

		if len(block.PrevHash) == 0 {
			break
			fmt.Printf("区块遍历完成退出!")
		}
	}

	return txs
}
