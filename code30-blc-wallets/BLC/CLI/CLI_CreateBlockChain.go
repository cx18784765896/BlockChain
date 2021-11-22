package CLI

import "BlockChain/code30-blc-wallets/BLC"

// 创建区块链
func (cli *CLI) creatBLC(address string) {
	BLC.CreatBlockchianWithGenesisBlock(address)
}