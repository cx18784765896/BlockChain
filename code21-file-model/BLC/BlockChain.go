package BLC

import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"strconv"
)

const dbName = "bc"
const blockTableName = "blocks"

// 基本的区块链结构
type BlockChian struct {
	//Blocks []*Block     // 存储有序的区块
	DB		 *bolt.DB     // 数据库
	Tip      []byte       // 最新区块的哈希值
}

// 判断数据库是否存在
func DbExits() bool {
	if _,err := os.Stat(dbName);os.IsNotExist(err){
		return false
	}
	return true
}

// 初始化区块链
func CreatBlockchianWithGenesisBlock(address string) *BlockChian{
	// 判断数据库是否存在
	if DbExits() {
		fmt.Println("创世区块已存在...")
		os.Exit(1)    // 退出
	}
	// 创建或者打开数据库
	db,err := bolt.Open(dbName,0600,nil)
	if err != nil {
		log.Panicf("open the db failed! %v\n",err)
	}
	//defer db.Close()
	var blockHash []byte // 需要存储到数据库中的区块哈希
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil == b{
			b,err = tx.CreateBucket([]byte(blockTableName))
			if nil != err{
				log.Panicf("create the bucket failed! %v\n",err)
			}
		}
		if nil != b {
			// 创建coinbase交易
			txCoinbase := NewCoinbaseTransaction(address)
			// 添加创世区块
			genesisBlock := CreatGenesisBlock([]*Transaction{txCoinbase})
			err = b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if nil != err {
				log.Panicf("put the data of genesisBlock to db failed! %v\n",err)
			}
		    // 存储最新区块的哈希
		    err = b.Put([]byte("1"),genesisBlock.Hash)
		    if nil != err {
		    	log.Panicf("put the hash of latest block to db failed! %v\n",err)
			}
			blockHash = genesisBlock.Hash
		}
		if nil != err {
			log.Panicf("update the db failed! %v\n",err)
		}
		return nil
	})
	return &BlockChian{db,blockHash}
}

// 添加新的区块到区块链中
func (bc *BlockChian) AddBlock(txs []*Transaction) {
	//newBlock := NewBlock(height,preBlockHash,data)
	//bc.Blocks = append(bc.Blocks,newBlock)

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		// 1. 获取数据表
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {  // 2. 确保数据表存在
			// 3.获取最新的区块数据
			// latestHash := b.Get([]byte("1"))
			blockBytes := b.Get(bc.Tip)
			latestBlock := Deserialize(blockBytes)
			// 4. 创建新区块
			newBlock := NewBlock(latestBlock.Height + 1,latestBlock.Hash,txs)
			// 5. 存入数据库
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if nil != err{
				log.Panicf("put the data of new block into db failed! %v\n",err)
			}
			// 6. 更新最新区块的哈希
			err = b.Put([]byte("1"),newBlock.Hash)
			if nil != err{
				log.Panicf("put the hash of latest block into db failed! %v\n",err)
			}
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if nil != err {
		log.Panicf("update the db failed! %v\n",err)
	}
}

// 输出区块链中的所有区块信息
func (bc *BlockChian) PrintChain() {
	fmt.Println("区块链完整信息：")
	var curBlock *Block
	//var curHash []byte = bc.Tip   // 获取最新区块哈希
	bcit := bc.Iterator()
	for {
		fmt.Printf("----------------------------------------------\n")
		curBlock  = bcit.Next()
		fmt.Printf("\t Height: %d \n",curBlock.Height)
		fmt.Printf("\t TimeStamp: %d \n",curBlock.TimeStamp)
		fmt.Printf("\t PreBlockHash: %x \n",curBlock.PreBlockHash)
		fmt.Printf("\t Hash: %x \n",curBlock.Hash)
		fmt.Printf("\t Trasactions: %v \n",curBlock.Txs)
		for _, tx := range curBlock.Txs{
			fmt.Printf("\t\t tx-hash: %x\n",tx.TxHash)
			for _, vin := range tx.Vins {
				fmt.Printf("\t\t vin-txhash: %x\n",vin.Txhash)
				fmt.Printf("\t\t vin-vout: %x\n",vin.Vout)
				fmt.Printf("\t\t vin-scriptSig: %x\n",vin.ScriptSig)
			}
			for _, vout := range tx.Vouts {
				fmt.Printf("\t\t vout-value: %d\n",vout.Value)
				fmt.Printf("\t\t vout-scriptPubkey: %s\n",vout.ScriptPubkey)
			}
		}
		fmt.Printf("\t Nonce: %d \n",curBlock.Nonce)

		// 3. 判断是否已经遍历到创世区块
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PreBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0{
			break  // 跳出循环
		}
	}
}

// 返回blockChain对象
func BlockChainObject() *BlockChian {
	// 读取数据库
	db , err := bolt.Open(dbName,0600,nil)
	if nil != err {
		log.Panicf("Open the db failed! %v\n",err)
	}
	var tip []byte   // 最新区块的哈希值
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			tip = b.Get([]byte("1"))
		}
		return nil
	})
	return &BlockChian{db,tip}
}

// 挖矿（打包交易，生成新的区块）
func (bc *BlockChian) MineNewBlock(from,to,amount []string)  {
	fmt.Printf("From:[%s]\n",from)
	fmt.Printf("To:[%s]\n",to)
	fmt.Printf("Amount:[%s]\n",amount)
	// 接收交易
	var txs []*Transaction  // 要打包的交易列表
	// 打包交易
	value,_ := strconv.Atoi(amount[0])
	tx := NewSimpleTransaction(from[0],to[0],value)
	txs = append(txs,tx)
	// 生成新的区块
	bc.AddBlock(txs)
}

// 查询指定地址的UTXO
func (bc *BlockChian) UnUTXOs(address string) []*UTXO {
	// 存储所有未花费输出
	var UnUTXOs []*UTXO
	// 存储所有已花费输出
	// key：与address相关的每个input所引用的交易hash
	// value：input引用交易的output的索引列表
	spentTxOutputs := make(map[string][]int)
	// 1.遍历区块链，查找与address相关的所有交易
	blockIterator := bc.Iterator()
	for {
		// 1.1 获取每一个区块信息
		block := blockIterator.Next()
		// 1.2 获取区块中的每一个交易信息
		for _,tx := range block.Txs{
			// 2. 先查找与address相关的输入
			if !tx.IsCoinbaseTransaction() {   //当该交易不是coinbase交易时
				// 获取输入信息
				for _,in := range tx.Vins {
					// 验证地址
					if in.UnlockWithAddress(address) {
						key := hex.EncodeToString(in.Txhash)
						spentTxOutputs[key] = append(spentTxOutputs[key],in.Vout)
					}
				}
			}
			// 3. 再查找与address相关的输出
			for index,vout := range tx.Vouts {
				// 地址验证
				if vout.UnlockScriptPubkeyWithAddress(address) {
					// 判断output是否已花费
					if len(spentTxOutputs) != 0 {   // 判断已花费输出是否为空
						for txhash,indexArray := range spentTxOutputs {   // 获取已花费输出中的每一条信息（引用的txhash,vout索引）
							for _,i := range indexArray{    // 获取引用的hash中的每一条vout索引信息
								// 判断该条输出是否在已花费输出中（txhash：是否是该条输出对应的交易，index：是否索引了该条输出）
								if txhash == hex.EncodeToString(tx.TxHash) && i == index{
									continue
								}else {
									utxo := &UTXO{
										Txhash: tx.TxHash,
										Index: index,
										Output: vout,
									}
									UnUTXOs = append(UnUTXOs,utxo)
								}
							}
						}
					} else {
						utxo := &UTXO{
							Txhash: tx.TxHash,
							Index: index,
							Output: vout,
						}
						UnUTXOs = append(UnUTXOs,utxo)
					}
				}
			}

		}
		// 4. 退出循环（判断是否循环到创世区块）
		var txhash big.Int
		txhash.SetBytes(block.PreBlockHash)
		if txhash.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return UnUTXOs
}

// 查询指定地址余额
func (bc *BlockChian) GetBalance(address string) int64 {
	utxos := bc.UnUTXOs(address)
	var amount int64
	for _,utxo := range utxos {
		amount += utxo.Output.Value
	}
	return amount
}