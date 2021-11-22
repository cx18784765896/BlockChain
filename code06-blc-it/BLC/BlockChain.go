package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
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

// 输出区块链中的所有区块信息
func (bc *BlockChian) PrintChain() {
	fmt.Println("区块链完整信息：")
	var curBlock *Block
	var curHash []byte = bc.Tip   // 获取最新区块哈希
	for {
		fmt.Printf("-------------------------------\n")
		bc.DB.View(func(tx *bolt.Tx) error {
			// 1. 获取表
			b := tx.Bucket([]byte(blockTableName))
			if nil != b {
				// 2. 获取当前区块信息
				blockBytes := b.Get(curHash)
				curBlock = Deserialize(blockBytes)
				/*
					TimeStamp 		int64       // 区块时间戳，区块产生的时间
					Height			int64	    // 区块高度（索引、号码），代表当前区块的高度
					PreBlockHash    []byte 		// 前一个区块（父区块）的hash值
					Hash			[]byte		// 当前区块的hash
					Data			[]byte		// 交易数据
					Nonce 			int64		// 用于工作量证明的随机数，碰撞次数
				 */
				fmt.Printf("\t Height: %d \n",curBlock.Height)
				fmt.Printf("\t TimeStamp: %d \n",curBlock.TimeStamp)
				fmt.Printf("\t PreBlockHash: %x \n",curBlock.PreBlockHash)
				fmt.Printf("\t Hash: %x \n",curBlock.Hash)
				fmt.Printf("\t Data: %s \n",string(curBlock.Data))
				fmt.Printf("\t Nonce: %d \n",curBlock.Nonce)
			}
			return nil
		})
		// 3. 判断是否已经遍历到创世区块
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0{
			break  // 跳出循环
		}
	    // 4. 更新curHash
	    curHash = curBlock.PreBlockHash
	}
}