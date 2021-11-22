package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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
func NewSimpleTransaction(from  string,to string,amount int,blockchain *BlockChian,txs []*Transaction) *Transaction {
	// 输入
	var txInputs []*TxInput
	// 输出
	var txOutputs []*TxOutput

	// 查找可用的UTXO
	money,spendableUTXODic := FindSpendableUTXO(from,int64(amount),blockchain,txs)
	for txHash,indexArray := range spendableUTXODic {
		txHashBytes,_ := hex.DecodeString(txHash)
		for _,index := range indexArray {
			txInput := &TxInput{txHashBytes,index,from}
			txInputs = append(txInputs,txInput)
		}
	}

	// 转账
	txOutput := &TxOutput{int64(amount),to}
	txOutputs = append(txOutputs,txOutput)

	//找零（目前数据写死了，逻辑上存在问题，因此输出数据是有问题的数据）
	txOutput = &TxOutput{money-int64(amount),from}
	txOutputs = append(txOutputs,txOutput)
	// 生成交易
	txSimple := &Transaction{nil,txInputs,txOutputs}
	//hash
	txSimple.HashTransaction()
	return txSimple
}

func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}