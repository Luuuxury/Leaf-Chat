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

	registdata := []byte(`{
		"UserRegist": {
			"RegistName": "admin-3",
			"RegistPW": "admin123-3"
		}
	}`)

	// 前两位为标识位, 第一个字节使用 1/0 表示所在字节后面还有/没有字节，第二个字节使用 1/0 表示所在字节后面有/没有字节
	buf := make([]byte, 2+len(registdata))
	binary.BigEndian.PutUint16(buf, uint16(len(registdata)))
	copy(buf[2:], registdata)
	conn.Write(buf)
}
