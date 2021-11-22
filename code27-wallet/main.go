package main

import (
	"BlockChain/code26-base58/BlockChain"
	"fmt"
)

func main() {
	msg := "this is the eg of base58 encode"
	//编码
	encoded := BlockChain.Base58Encode([]byte(msg))
	fmt.Printf("encoded:%s\n",encoded)

	// 解码
	decode_data := BlockChain.Base58Decode(encoded)
	fmt.Printf("msg:%v\n",string(decode_data))
}
