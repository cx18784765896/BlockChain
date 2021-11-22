package BLC

import "bytes"

// 交易输出
type TxOutput struct {
	// 有多少钱（金额）
	Value int64
	//// 钱是谁的（用户）
	//ScriptPubkey string
	// 公钥哈希
	PublicKeyHash []byte
}

// Output身份验证
func (out *TxOutput) UnlockScriptPubkeyWithAddress(address string) bool {
	pubHash := lock(address)
	return bytes.Compare(out.PublicKeyHash,pubHash) == 0
}

// 获取公钥哈希
func lock(address string) []byte {
	addrToBytes := Base58Decode([]byte(address))
	pubKeyHash := addrToBytes[1:len(addrToBytes)-checkSumLen]
	return pubKeyHash
}

// 生成一个output对象
func NewOutput(value int64,address string) *TxOutput {
	txOutput := &TxOutput{}
	txOutput.Value = value
	txOutput.PublicKeyHash = lock(address)
	return txOutput
}