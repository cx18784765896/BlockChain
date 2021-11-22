package BlockChain
// 钱包合集相关文件

// 钱包合集结构
type Wallets struct {
	Wallets map[string] *Wallet
}

// 初始化一个钱包合集
func InitWallets() *Wallets {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string] *Wallet)
	return wallets
}

// 新建钱包并将其添加到钱包合集中
func (wallets *Wallets) CreatWallet() {
	wallet := NewWallet()  // 新建钱包对象
	wallets.Wallets[string(wallet.GetAddress())] = wallet
}