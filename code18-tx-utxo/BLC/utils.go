package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

// 整数转字节
func IntToHex(data int64) []byte {
	buffer := new (bytes.Buffer)  // 新建一个buffer
	err := binary.Write(buffer,binary.BigEndian,data)
	if nil != err {
		log.Panicln("int to []byte failed! %v\n",err)
	}
	return buffer.Bytes()
}
// Windows下标准JSON的输入格式
/*
	bc.exe send -from "[\"Alice\",\"Bob\"]" -to "[\"Bob\",\"Sirius\"]" -amount "[\"5\",\"6\"]"
*/
// 标准JSON格式转数组
func JSONToArray(jsonString string) []string {
	var strArr []string
	// Json.Unmarshal
	if err := json.Unmarshal([]byte(jsonString),&strArr);err != nil {
		log.Panicf("Json to []string failed! %v\n",err)
	}
	return strArr
}