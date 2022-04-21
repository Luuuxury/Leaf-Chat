package main

import (
	"encoding/binary"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:21002")
	if err != nil {
		panic(err)
	}

	logindata := []byte(`{
	"UserLogin":{
		"LoginName": "admin-2",
		"LoginPW": "admin123-2"
		}
	}`)

	// 前两位为标识位: 第一个字节使用 1/0 表示所在字节后面还有/没有字节，第二个字节使用 1/0 表示所在字节后面有/没有字节
	writeBuf := make([]byte, 2+len(logindata))
	binary.BigEndian.PutUint16(writeBuf, uint16(len(logindata)))
	copy(writeBuf[2:], logindata)
	conn.Write(writeBuf)
}
