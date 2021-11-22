package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI结构
type CLI struct {
	BC *BlockChian
}

// 展示用法
func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Printf("\t createblockchain -- 创建区块链 \n")
	fmt.Printf("\t addblock -data DATA --添加区块 \n")
	fmt.Printf("\t printchain --输出区块链信息 \n")
}

// 校验，如果只输入了程序命令，就输出指令用法并且退出程序
func IsValidArgs()  {
	if len(os.Args) < 2 {
		PrintUsage() //打印用法
		os.Exit(1)  // 退出程序
	}
}

// 添加区块
func (cli *CLI) addBlock(data string)  {
	cli.BC.AddBlock([]byte(data))
}
// 输出区块链信息
func (cli *CLI) printBLC() {
	cli.BC.PrintChain()
}
// 创建区块链
func (cli *CLI) creatBLC() {
	CreatBlockchianWithGenesisBlock()
}

// 运行函数
func (cli *CLI) Run() {
	// 1. 检测参数数量
	IsValidArgs()
	// 2. 新建命令
	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	createBlcCmd := flag.NewFlagSet("createblockchain",flag.ExitOnError)
	// 3. 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data","send 100 BTC to SZC","交易数据")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if nil != err{
			log.Panicf("prase cmd addblock failed! %v\n",err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if nil != err{
			log.Panicf("prase cmd printchain failed! %v\n",err)
		}
	case "createblockchain":
		err := createBlcCmd.Parse(os.Args[2:])
		if nil != err{
			log.Panicf("prase cmd createblockchain failed! %v\n",err)
		}
	default:
		PrintUsage()
		os.Exit(1)
	}

	// 添加区块
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockArg)
	}

	// 输出区块链信息
	if printChainCmd.Parsed() {
		cli.printBLC()
	}

	// 创建区块链
	if createBlcCmd.Parsed() {
		cli.creatBLC()
	}

}