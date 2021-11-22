package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targitBit = 16

// 工作量证明
type ProofOfWork struct {
	Block 	 *Block     // 对指定区块进行验证
	target   *big.Int	// 大数据存储
}

// 创建新的POW对象，即将生成难度比较时的target和区块block放到一个POW对象中
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 假设数据（1）总长8位，targetBit等于2，则8-2=6，则targit等于1左移6位
	// 1<<6 = 0010 0000 =64，则target=64
	// 对此处计算target同理，但sha256计算出来的数据是256位
	target = target.Lsh(target,256 - targitBit)
	return &ProofOfWork{block,target}
}

// 开始工作量证明
func (proofOfWork *ProofOfWork) Run() ([]byte,int64) {
	var nonce = 0      // 碰撞次数
	var hash [32]byte  // 生成的哈希值
	var hashInt big.Int   // 用于存储哈希转换之后生成的数据，最终和target 数据进行比较
	for {
		// 数据拼接
		dataBytes := proofOfWork.PrepareData(nonce)
		hash = sha256.Sum256(dataBytes)
		hashInt.SetBytes(hash[:])
		fmt.Printf("hash: \r%x",hash)   // 以十六进制输出
		// 难度比较
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	fmt.Printf("\n碰撞次数：%d\n",nonce)
	return hash[:],int64(nonce)
}

// 准备数据，将区块链属性链接起来，返回一个字节数组
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PreBlockHash,
		pow.Block.HashTransaction(),
		IntToHex(pow.Block.TimeStamp),
		IntToHex(pow.Block.Height),
		IntToHex(int64(nonce)),
		IntToHex(targitBit),
	},[]byte{})
	return data
}