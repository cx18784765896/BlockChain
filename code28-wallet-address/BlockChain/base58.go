package BlockChain

import (
	"bytes"
	"math/big"
)

// 实现base58编码

//base58字符表
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// 编码函数
func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)    // Bytes转换为bigInt
	//fmt.Printf("x:%v\n",x)
	base := big.NewInt(int64(len(b58Alphabet)))   // 设置一个base58求模的基数
	zero := big.NewInt(0)
	mod := &big.Int{}   // 余数
	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)   // x除以base余mod
		// 以余数为下标，取值
		result = append(result,b58Alphabet[mod.Int64()])
	}
	result = Reverse(result)
	//fmt.Printf("result: %s\n",result)

	for b:=range input {   // b代表输入的切片的下标
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]},result...)
		} else {
			break
		}
	}
	return result
}
// 解码函数
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for b:= range input {
		if b == 0x00 {
			zeroBytes++
		}
	}
	data := input[zeroBytes:]
	for _,b := range data {
		// 获取bytes数组中指定数字第一次出现的索引
		charIndex := bytes.IndexByte(b58Alphabet,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}
	decode := result.Bytes()
	decode = append(bytes.Repeat([]byte{byte(0x00)},zeroBytes),decode...)
	return decode
}