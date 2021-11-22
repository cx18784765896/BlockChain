package CLI

import (
	"BlockChain/code31-wallet-in-out/BLC"
	"fmt"
)

// 获取钱包合集中的所有地址
func (cli *CLI) GetWalletsAddr()  {
	fmt.Println("打印钱包中的所有地址：")
	wallets,_ := BLC.InitWallets()
	for address,_ := range wallets.Wallets {
		fmt.Printf("address: [%s] \n",address)
	}
}
