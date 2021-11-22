package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 交易相关

type Transaction struct {
	// tx hash （交易的唯一标识）
	TxHash []byte
	// 输入
	Vins    []*TxInput
	// 输出
	Vouts	[]*TxOutput
}

// 生成交易hash
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panicf("tx hash generate failed! %v\n")
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

// 生成coinbase交易
func NewCoinbaseTransaction(address string) *Transaction {
	// 输入
	txInput := &TxInput{[]byte{},-1,"Genesis Data"}
	// 输出
	txOutput := &TxOutput{10,address}
    // 生成交易
	txCoinbase := &Transaction{nil,[]*TxInput{txInput},[]*TxOutput{txOutput}}
	// hash
	txCoinbase.HashTransaction()
	return txCoinbase
}

// 生成普通转账交易