package CLI

import "BlockChain/code21-file-model/BLC"

// 创建区块链
func (cli *CLI) creatBLC(address string) {
	BLC.CreatBlockchianWithGenesisBlock(address)
}