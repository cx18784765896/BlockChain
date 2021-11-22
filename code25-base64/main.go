package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Sirius"
	// 编码
	encode := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encode)

	// 解码
	decode ,err := base64.StdEncoding.DecodeString("U2lyaXVz")
	if nil!=err {
		panic(err)
	}
	fmt.Println(string(decode))
}
