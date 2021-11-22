package BLC
// 交易相关

type Transaction struct {
	// tx hash （交易的唯一标识）
	TxHash []byte

}

// 生成coinbase交易

// 生成普通转账交易