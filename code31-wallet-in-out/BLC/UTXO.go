package BLC

type UTXO struct {
	// 对应的交易哈希
	Txhash []byte
	// 对应的输出的索引
	Index  int
	// 对应的输出
	Output *TxOutput
}