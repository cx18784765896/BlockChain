package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToHex(data int64) []byte {
	buffer := new (bytes.Buffer)  // 新建一个buffer
	err := binary.Write(buffer,binary.BigEndian,data)
	if nil != err {
		log.Panicln("int to []byte failed! %v\n",err)
	}
	return buffer.Bytes()
}