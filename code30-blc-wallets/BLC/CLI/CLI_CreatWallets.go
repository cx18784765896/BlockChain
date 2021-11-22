package CLI

import (
	"BlockChain/code30-blc-wallets/BLC"
	"fmt"
)

// 创建钱包集合
func (cli *CLI) CreateWallets() {
	wallets,_ := BLC.InitWallets()
	wallets.CreatWallet()
	fmt.Printf("wallets:%v\n",wallets)
}
