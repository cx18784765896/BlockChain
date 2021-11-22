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
	fmt.Printf("\t createblockchain -address Address -- 创建区块链 \n")
	fmt.Printf("\t addblock -data DATA --添加区块 \n")
	fmt.Printf("\t printchain --输出区块链信息 \n")
	fmt.Printf("\t send -from FROM_ADDR -to TO_ADDR -amount AMOUNT --发起交易 \n")
	fmt.Printf("\t getbalance -address Address --查询余额 \n")
}

// 校验，如果只输入了程序命令，就输出指令用法并且退出程序
func IsValidArgs()  {
	if len(os.Args) < 2 {
		PrintUsage() //打印用法
		os.Exit(1)  // 退出程序
	}
}

// 查询余额
func (cli *CLI) getBalance(from string) {
	if dbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BlockChainObject()  // 获取区块链对象
	defer blockChain.DB.Close()
	blockChain.UnUTXOs(from)
}

// 发起交易
func (cli *CLI) send(from,to,amonut []string) {
	if dbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BlockChainObject()  // 获取区块链对象
	defer blockChain.DB.Close()
	blockChain.MineNewBlock(from,to,amonut)
}

// 添加区块
func (cli *CLI) addBlock(txs []*Transaction)  {
	if dbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BlockChainObject()  // 获取区块链对象
	defer blockChain.DB.Close()
	blockChain.AddBlock([]*Transaction{})
}
// 输出区块链信息
func (cli *CLI) printBLC() {
	if dbExits() == false {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	blockChain := BlockChainObject()  // 获取区块链对象
	defer blockChain.DB.Close()
	blockChain.PrintChain()
}
// 创建区块链
func (cli *CLI) creatBLC(address string) {
	CreatBlockchianWithGenesisBlock(address)
}

// 运行函数
func (cli *CLI) Run() {
	// 1. 检测参数数量
	IsValidArgs()
	// 2. 新建命令
	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	createBlcCmd := flag.NewFlagSet("createblockchain",flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send",flag.ExitOnError)
	getBlanceCmd := flag.NewFlagSet("getbalance",flag.ExitOnError)
	// 3. 获取命令行参数
	flagAddBlockArg := addBlockCmd.String("data","send 100 BTC to SZC","交易数据")
	flagCreateBlcArg := createBlcCmd.String("address","","地址")
	flagFromArg := sendCmd.String("from","","源地址")
	flagToArg := sendCmd.String("to","","目标地址")
	flagAmount := sendCmd.String("amount","","转账金额")
	flagBalanceArg := getBlanceCmd.String("address","","查询地址")
	switch os.Args[1] {
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if nil != err{
			log.Panicf("prase cmd send failed! %v\n",err)
		}
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
	case "getbalance":
		err := getBlanceCmd.Parse(os.Args[2:])
		if nil != err{
			log.Panicf("prase cmd getbalance failed! %v\n",err)
		}
	default:
		PrintUsage()
		os.Exit(1)
	}

	// 查询余额
	if getBlanceCmd.Parsed() {
		if *flagBalanceArg == "" {
			fmt.Println("未指定查询地址！")
			PrintUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagBalanceArg)
	}

	// 发起交易
	if sendCmd.Parsed() {
		if *flagFromArg == "" {
			fmt.Println("源地址不能为空！")
			PrintUsage()
			os.Exit(1)
		}
		if *flagToArg == "" {
			fmt.Println("目标地址不能为空！")
			PrintUsage()
			os.Exit(1)
		}
		if *flagAmount == "" {
			fmt.Println("转账金额不能为空！")
			PrintUsage()
			os.Exit(1)
		}
		cli.send(JSONToArray(*flagFromArg),JSONToArray(*flagToArg),JSONToArray(*flagAmount))
	}

	// 添加区块
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock([]*Transaction{})
	}

	// 输出区块链信息
	if printChainCmd.Parsed() {
		cli.printBLC()
	}

	// 创建区块链
	if createBlcCmd.Parsed() {
		if *flagCreateBlcArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.creatBLC(*flagCreateBlcArg)
	}

}