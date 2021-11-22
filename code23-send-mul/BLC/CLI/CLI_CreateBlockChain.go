package CLI

import "BlockChain/code23-send-mul/BLC"

// 创建区块链
func (cli *CLI) creatBLC(address string) {
	BLC.CreatBlockchianWithGenesisBlock(address)
}