package CLI

import "BlockChain/code22-send-utxo/BLC"

// 创建区块链
func (cli *CLI) creatBLC(address string) {
	BLC.CreatBlockchianWithGenesisBlock(address)
}