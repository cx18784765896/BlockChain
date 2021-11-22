package main

import (
	"BlockChain/code18-tx-utxo/BLC"

)

// 测试
func main()  {
	//blockChain := BLC.CreatBlockchianWithGenesisBlock()
	//defer blockChain.DB.Close()
	//blockChain.AddBlock([]byte("szc"))
	//blockChain.AddBlock([]byte("love"))
	//blockChain.AddBlock([]byte("cx"))
	cli := BLC.CLI{}
	cli.Run()
}
