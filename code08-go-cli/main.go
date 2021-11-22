package main

import (
	"flag"
	"fmt"
	"net/url"
)

//// 命令行操作
//func main() {
//	args := os.Args
//	fmt.Printf("args[0]:%v\n",args[0])
//	if len(args) >= 2 {
//		for i := 0; i < len(args); i++ {
//			fmt.Printf("args[%d]:%v\n",i,args[i])
//		}
//	}
//}

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func main() {
	//fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	//fs.Var(&URLValue{u}, "url", "URL to parse")
	//
	//fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	//fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

	flagPrintChain := flag.String("printChain","cx","print the info")
	flag.Parse()  // 解析
	fmt.Printf("the flag of string : %v\n",*flagPrintChain)
}
