package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

// 钱包相关
// 版本
const version = byte(0x00)
//检验和长度
const checkSumLen = 4

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

// 生成地址
func (w *Wallet) GetAddress() []byte {
	// 1. 获取公钥哈希
	pubKeyHash := PubKeyHash(w.PublicKey)
	// 2. 生成version并拼接到公钥哈希中
	version_pubKeyHash := append([]byte{version},pubKeyHash...)
	// 3. 生成校验和
	checkSum :=CheckSum(version_pubKeyHash)
	// 4. 将检验和拼接到version+公钥hash中
	bytes := append(version_pubKeyHash,checkSum...)
	// 5. 调用base58生成地址
	address := Base58Encode(bytes)

	return address
}

// 生成公钥哈希
func PubKeyHash(data []byte) []byte {
	// 1. 对公钥进行sha256哈希
	hash256 := sha256.New()
	hash256.Write(data)
	hash := hash256.Sum(nil)

	// 2. 对上一步生成的哈希做Ripemd160哈希
	hash160 := ripemd160.New()
	hash160.Write(hash)
	hashRipemd := hash160.Sum(nil)

	return hashRipemd
}

// 生成校验和
func CheckSum(data []byte) []byte {
	// 1. 对公钥哈希进行第一次sha256
	hash1 := sha256.Sum256(data)

	// 2. 对公钥哈希进行第二次sha256
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:checkSumLen]
}

// 地址有效性检验
func IsVaildForAddress(address []byte) bool {
	// 1. 解码
	decoded := Base58Decode(address)  //25位
	// 2. 获取checkSum
	checkSumBytes := decoded[len(decoded)-checkSumLen:]
	// 3. 获取version_pubKeyHash
	version_pubKeyHash := decoded[:len(decoded)-checkSumLen]
	// 4. 验证（重新计算version_pubKeyHash的checkSum是否等于获取的checkSum）
	checkSum := CheckSum(version_pubKeyHash)
	if bytes.Compare(checkSum,checkSumBytes) == 0 {
		return true
	}
	return false
}