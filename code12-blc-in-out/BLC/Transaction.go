package BLC
// 交易相关

type Transaction struct {
	// tx hash （交易的唯一标识）
	TxHash []byte
	// 输入
	Vins    []*TxInput
	// 输出
	Vouts	[]*TxOutput
}

// 生成coinbase交易

// 生成普通转账交易