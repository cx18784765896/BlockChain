package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbName = "bc"
const blockTableName = "blocks"

// 基本的区块链结构
type BlockChian struct {
	//Blocks []*Block  // 存储有序的区块
	DB		 *bolt.DB     // 数据库
	Tip      []byte       // 最新区块的哈希值
}

// 初始化区块链
func CreatBlockchianWithGenesisBlock() *BlockChian{
	// 创建或者打开数据库
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Panicf("open the db failed! %v\n",err)
	}
	//defer db.Close()
	var blockHash []byte // 需要存储到数据库中的区块哈希
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil == b{
			b,err = tx.CreateBucket([]byte(blockTableName))
			if nil != err{
				log.Panicf("create the bucket failed! %v\n",err)
			}
		}
		if nil != b {
			// 添加创世区块
			genesisBlock := CreatGenesisBlock("the init of block chain")
			err = b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if nil != err {
				log.Panicf("put the data of genesisBlock to db failed! %v\n",err)
			}
		    // 存储最新区块的哈希
		    err = b.Put([]byte("1"),genesisBlock.Hash)
		    if nil != err {
		    	log.Panicf("put the hash of latest block to db failed! %v\n",err)
			}
			blockHash = genesisBlock.Hash
		}
		if nil != err {
			log.Panicf("update the db failed! %v\n",err)
		}
		return nil
	})
	return &BlockChian{db,blockHash}
}

// 添加新的区块到区块链中
func (bc *BlockChian) AddBlock(data []byte) {
	//newBlock := NewBlock(height,preBlockHash,data)
	//bc.Blocks = append(bc.Blocks,newBlock)

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		// 1. 获取数据表
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {  // 2. 确保数据表存在
			// 3.获取最新的区块数据
			// latestHash := b.Get([]byte("1"))
			blockBytes := b.Get(bc.Tip)
			latestBlock := Deserialize(blockBytes)
			// 4. 创建新区块
			newBlock := NewBlock(latestBlock.Height + 1,latestBlock.Hash,data)
			// 5. 存入数据库
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if nil != err{
				log.Panicf("put the data of new block into db failed! %v\n",err)
			}
			// 6. 更新最新区块的哈希
			err = b.Put([]byte("1"),newBlock.Hash)
			if nil != err{
				log.Panicf("put the hash of latest block into db failed! %v\n",err)
			}
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if nil != err {
		log.Panicf("update the db failed! %v\n",err)
	}
}