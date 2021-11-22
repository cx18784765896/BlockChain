package CLI

import (
	"BlockChain/code21-file-model/BLC"
	"fmt"
	"os"
)

// 输出区块链信息
func (cli *CLI) printBLC() {
	if BLC.DbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BLC.BlockChainObject() // 获取区块链对象
	defer blockChain.DB.Close()
	blockChain.PrintChain()
}