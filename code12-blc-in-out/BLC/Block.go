package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// 实现一个最基本的区块结构
type Block struct {
	TimeStamp 		int64      	   	 // 区块时间戳，区块产生的时间
	Height			int64	    	 // 区块高度（索引、号码），代表当前区块的高度
	PreBlockHash    []byte 			 // 前一个区块（父区块）的hash值
	Hash			[]byte			 // 当前区块的hash
	//Data			[]byte			 // 交易数据
	Txs 			[]*Transaction   // 交易
	Nonce 			int64			 // 用于工作量证明的随机数，碰撞次数
}

// 创建一个区块
func NewBlock(height int64,preBlockHsh []byte,txs []*Transaction) *Block{
	var block Block
	block = Block{
		TimeStamp: time.Now().Unix(),
		Height: height,
		PreBlockHash: preBlockHsh,
		Txs: txs,
	}
	//block.SetHash()
	//生成当前区块hash
	pow := NewProofOfWork(&block)   // 生成一个计算POW的对象pow（包含的数据有传入的区块block，和计算出来的target）
	hash,nonce := pow.Run()        // 调用pow的方法Run，生成hash和nonce
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

// 生成创世区块
func CreatGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(0,nil,[]*Transaction{})
}

// 序列化，将区块数据序列化为[]byte（字节数组）
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)  // 新建encode对象
	if err := encoder.Encode(block);nil != err {
		log.Panicf("serialize the block to byte failed! %v\n",err)
	}
	return result.Bytes()
}

// 反序列化，将字节数组结构化为区块
func Deserialize(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block);nil != err {
		log.Panicf("deserialize the []byte to block failed! %v\n",err)
	}
	return &block
}

// 把区块中的所有交易结构转换成[]byte
func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	for _,tx := range block.Txs {
		txHashes = append(txHashes,tx.TxHash)
	}
	// sha256
	txHash := sha256.Sum256(bytes.Join(txHashes,[]byte{}))
	return txHash[:]
}