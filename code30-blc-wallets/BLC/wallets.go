package BLC

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

// 钱包合集相关文件

// 钱包信息文件
const walletFile = "wallets.dat"

// 钱包合集结构
type Wallets struct {
	Wallets map[string] *Wallet
}

// 初始化一个钱包合集
func InitWallets() (*Wallets,error) {
	// 1. 判断文件是否存在
	if _,err := os.Stat(walletFile);os.IsNotExist(err) {
		// 若不存在，则创建文件
		wallets := &Wallets{}
		wallets.Wallets = make(map[string] *Wallet)
		return wallets, err
	}
	// 2. 文件存在，读取内容
	fileContent,err := ioutil.ReadFile(walletFile)
	if nil != err {
		log.Panicf("read file content failed! %v\n",err)
	}
	var wallets Wallets
	// register 适用于需要解析的参数中包含interface
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if nil != err {
		log.Panicf("decode file content failed ! %v\n",err)
	}
	return &wallets,nil
}

// 新建钱包并将其添加到钱包合集中
func (wallets *Wallets) CreatWallet() {
	wallet := NewWallet()  // 新建钱包对象
	wallets.Wallets[string(wallet.GetAddress())] = wallet
	// 把钱包存储到文件中
	wallets.SaveWallets()
}

// 持久化钱包信息
func (wallets *Wallets) SaveWallets() {
	var content bytes.Buffer
	// 注册
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	// 序列化钱包数据
	err := encoder.Encode(&wallets)
	if nil != err {
		log.Panicf("encode the construct of wallet failed ! %v\n",err)
	}
	// 清空文件再去存储（此处只保存一条数据，但该条数据会保存到目前为止所有地址的集合）
	err = ioutil.WriteFile(walletFile,content.Bytes(),0644)
	if nil != err {
		log.Panicf("write the content of wallets to file [%s] failed ! %v\n",walletFile,err)
	}
}