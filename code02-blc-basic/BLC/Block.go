package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

// 实现一个最基本的区块结构
type Block struct {
	TimeStamp 		int64       // 区块时间戳，区块产生的时间
	Height			int64	    // 区块高度（索引、号码），代表当前区块的高度
	PreBlockHash    []byte 		// 前一个区块（父区块）的hash值
	Hash			[]byte		// 当前区块的hash
	Data			[]byte		// 交易数据
}

// 创建一个区块
func NewBlock(height int64,preBlockHsh []byte,data []byte) *Block{
	var block Block
	block = Block{
		TimeStamp: time.Now().Unix(),
		Height: height,
		PreBlockHash: preBlockHsh,
		Data: data,
	}
	block.SetHash()  // 生成当前区块hash
	return &block
}

// 计算区块hash
func (b *Block) SetHash() {
	// 整数型转字节型
	heightBytes := IntToHex(b.Height)
	timeStampBytes := IntToHex(b.TimeStamp)
	// 拼接所有属性，进行哈希
	blockBytes := bytes.Join([][]byte{heightBytes,timeStampBytes,b.PreBlockHash,b.Data},[]byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]  // 为什么要用[:]而不是直接用[]?
}

// 生成创世区块
func CreatGenesisBlock(data string) *Block {
	return NewBlock(0,nil,[]byte(data))
}