package main

import (
	"BlockChain/code05-blc-db/BLC"
	"fmt"
	"github.com/boltdb/bolt"
)

// 测试
func main()  {
	////block := BLC.NewBlock(1,nil,[]byte("this is the basic concept block"))
	////block := BLC.CreatGenesisBlock("this is the basic concept block")
	//blockChain := BLC.CreatBlockchianWithGenesisBlock()
	//fmt.Printf("blockchain: %v\n",blockChain)
	//blockChain.AddBlock(
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	[]byte("Ami send 100 BTC to Bob"),
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlock(
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	[]byte("Ami send 35 BTC to Bob"),
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//blockChain.AddBlock(
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Height+1,
	//	[]byte("Ami send 20 BTC to Bob"),
	//	blockChain.Blocks[len(blockChain.Blocks)-1].Hash)
	//length := len(blockChain.Blocks)
	//fmt.Printf("the length of blockChain is %d\n",length)
	//for i := 0; i < length; i++ {
	//	fmt.Printf("the %d th block is %v\n",i,blockChain.Blocks[i])
	//}
	////fmt.Printf("block: %v\n",block)

	//db,err := bolt.Open("cb",0600,nil)
	//if err != nil {
	//	log.Panicf("open the db failed! %v\n",err)
	//}
	//defer db.Close()
	//genesisBlock := BLC.CreatGenesisBlock("genesisBlock")
	//db.Update(func(tx *bolt.Tx) error {
	//	b,err := tx.CreateBucket([]byte("blocks"))
	//	if nil != err{
	//		log.Panicf("create the bucket failed! %v\n",err)
	//	}
	//	blockData := genesisBlock.Serialize()
	//	err = b.Put([]byte("1"),blockData)
	//	if nil != err {
	//		log.Panicf("put the data to db failed! %v\n",err)
	//	}
	//	return nil
	//})
	//err = db.View(func(tx *bolt.Tx) error {
	//	b := tx.Bucket([]byte("blocks"))
	//	if nil != b {
	//		data := b.Get([]byte("1"))
	//		fmt.Printf("data: %v\n",genesisBlock.Deserialize(data))
	//	}
	//	return nil
	//})
	//if nil != err {
	//	log.Panicf("get the data of block failed! %v\n",err)
	//}
	blockChain := BLC.CreatBlockchianWithGenesisBlock()
	defer blockChain.DB.Close()
	blockChain.AddBlock([]byte("szc"))
	blockChain.AddBlock([]byte("love"))
	blockChain.AddBlock([]byte("cx"))
	blockChain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		if nil != b {
			latestHash := b.Get([]byte("1"))
			fmt.Printf("the latest blockHash: %x\n",latestHash)
			block := b.Get(latestHash)
			fmt.Printf("the height of blockChain: %d\n",BLC.Deserialize(block).Height)
		} else {
			fmt.Printf("the bucket is nil !\n")
		}
		return nil
	})
}
