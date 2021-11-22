package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

// 基本的迭代器结构
type BlockChainIterator struct {
	DB   		 *bolt.DB       // 数据库
	CurrentHash  []byte			// 当前区块的哈希值
}

// 创建一个迭代器对象
func (bc *BlockChian) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.DB,bc.Tip}
}

// 遍历
func (bcit *BlockChainIterator) Next() *Block {
	var block *Block
	err := bcit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b{
			// 获取指定区块的哈希
			currentBlockBytes := b.Get(bcit.CurrentHash)
			block = Deserialize(currentBlockBytes)
			// 更新迭代器中的currentHash
			bcit.CurrentHash = block.PreBlockHash
		}
		return nil
	})
	if nil != err{
		log.Panicf("iterator the db of blockchain failed! %v\n",err)
	}
	return block
}