package BLC
// 交易输入
type TxInput struct {
	// 交易哈希（不是当前交易的哈希）
	Txhash  []byte
	// 引用的上一笔交易的output的索引
	Vout  int
	// 用户名
	ScriptSig  string
}

func (in *TxInput) UnlockWithAddress(address string) bool {
	return in.ScriptSig == address
}