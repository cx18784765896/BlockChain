package BLC

// 基本的区块链结构
type BlockChian struct {
	Blocks []*Block  // 存储有序的区块
}

// 初始化区块链
func CreatBlockchianWithGenesisBlock() *BlockChian{
	// 添加创世区块
	genesisBlock := CreatGenesisBlock("the init of block chain")
	return &BlockChian{[]*Block{genesisBlock}}
}

// 添加新的区块到区块链中
func (bc *BlockChian) AddBlock(height int64,data []byte,preBlockHash []byte) {
	newBlock := NewBlock(height,preBlockHash,data)
	bc.Blocks = append(bc.Blocks,newBlock)
}