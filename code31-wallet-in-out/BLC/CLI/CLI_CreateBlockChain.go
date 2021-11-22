package CLI

import "BlockChain/code31-wallet-in-out/BLC"

// 创建区块链
func (cli *CLI) creatBLC(address string) {
	BLC.CreatBlockchianWithGenesisBlock(address)
}