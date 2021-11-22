package main

import (
	"blockchain/code02-blc-basic/BLC"
	"fmt"
)

// 测试
func main()  {
	//block := BLC.NewBlock(1,nil,[]byte("this is the basic concept block"))
	//block := BLC.CreatGenesisBlock("this is the basic concept block")
	blockChain := BLC.CreatBlockchianWithGenesisBlock()
	fmt.Printf("blockchain: %v\n",blockChain)
	blockChain.AddBlock(
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		[]byte("Ami send 100 BTC to Bob"),
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlock(
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		[]byte("Ami send 35 BTC to Bob"),
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	blockChain.AddBlock(
		blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
		[]byte("Ami send 20 BTC to Bob"),
		blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	length := len(blockChain.Blocks)
	fmt.Printf("the length of blockChain is %d\n",length)
	for i := 0; i < length; i++ {
		fmt.Printf("the %d th block is %v\n",i,blockChain.Blocks[i])
	}
	//fmt.Printf("block: %v\n",block)
}
