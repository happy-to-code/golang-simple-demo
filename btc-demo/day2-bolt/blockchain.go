package main

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
