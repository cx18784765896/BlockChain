package main

import (
	"blockchain/code01-bc-basic/BLC"
	"fmt"
)

// 测试
func main()  {
	block := BLC.NewBlock(1,nil,[]byte("this is the basic concept block"))
	fmt.Printf("block: %v\n",block)
}
