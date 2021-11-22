package BLC

import "bytes"

// 交易输入
type TxInput struct {
	// 交易哈希（不是当前交易的哈希）
	TxHash []byte
	// 引用的上一笔交易的output的索引
	Vout  int
	// 签名
	Signature []byte
	// 公钥
	PublicKey []byte

}

// Input身份验证
func (in *TxInput) UnlockPublicKeyHash(publicKeyHash []byte) bool {
	// 获取input的Ripemd160Hash(公钥哈希)
	inputPublicKeyHash := PubKeyHash(in.PublicKey)
	return bytes.Compare(publicKeyHash,inputPublicKeyHash) ==0
}