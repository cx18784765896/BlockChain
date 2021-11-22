package main

import (
	"BlockChain/code28-wallet-address/BlockChain"
	"fmt"
)

func main() {
	// 1. 生成钱包
	wallet := BlockChain.NewWallet()
	// 2. 生成地址
	address := wallet.GetAddress()
	fmt.Printf("address: %s\n",address)
	fmt.Printf("the validation of address: %s is %v !\n",address,BlockChain.IsVaildForAddress(address))
}
