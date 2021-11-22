package CLI

import (
	"BlockChain/code22-send-utxo/BLC"
	"fmt"
	"os"
)

// 查询余额
func (cli *CLI) getBalance(from string) {
	if BLC.DbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BLC.BlockChainObject() // 获取区块链对象
	defer blockChain.DB.Close()
	amount := blockChain.GetBalance(from)
	fmt.Printf("\t 地址：%v 的余额为：%d \n",from,amount)
}