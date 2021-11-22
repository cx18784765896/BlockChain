package CLI

import (
	"BlockChain/code23-send-mul/BLC"
	"fmt"
	"os"
)

// 发起交易
func (cli *CLI) send(from,to,amonut []string) {
	if BLC.DbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BLC.BlockChainObject() // 获取区块链对象
	defer blockChain.DB.Close()
	blockChain.MineNewBlock(from,to,amonut)
}