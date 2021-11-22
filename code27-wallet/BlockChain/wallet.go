package BlockChain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

// 钱包相关

// 钱包结构
type Wallet struct {
	// 私钥
	PrivateKey ecdsa.PrivateKey
	// 公钥
	PublicKey []byte
}

// 新建钱包
func NewWallet() *Wallet {
	privateKey,publicKey := newKeyPair()
	return &Wallet{PrivateKey: privateKey,PublicKey: publicKey}
}

// 新建公钥-私钥对
func newKeyPair() (ecdsa.PrivateKey,[]byte) {
	// 生成椭圆参数
	curve := elliptic.P256()
	// 生成密钥对
	priv,err := ecdsa.GenerateKey(curve,rand.Reader)
	if nil !=err {
		log.Panicf("ecdsa generate key failed! %v\n",err)
	}
	pubKey := append(priv.PublicKey.X.Bytes(),priv.PublicKey.Y.Bytes()...)
	return *priv,pubKey
}
