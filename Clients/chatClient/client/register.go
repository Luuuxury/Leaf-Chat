package main

import (
	"encoding/binary"
	"fmt"
	"leaf-chat/Servers/msg"
	"net"
)

func init() {
	msg.Processor.SetHandler(&msg.RegistResult{}, handleCheckRegist)
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:21002")
	if err != nil {
		panic(err)
	}

	registdata := []byte(`{
		"UserRegist": {
			"RegistName": "admin-9",
			"RegistPW": "admin123-9"
		}
	}`)

	// 前两位为标识位, 第一个字节使用 1/0 表示所在字节后面还有/没有字节，第二个字节使用 1/0 表示所在字节后面有/没有字节
	buf := make([]byte, 2+len(registdata))
	binary.BigEndian.PutUint16(buf, uint16(len(registdata)))
	copy(buf[2:], registdata)
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("客户端写入数据出错了")
	}

}

func handleCheckRegist(args []interface{}) {

	recv := args[0].(*msg.RegistResult)
	fmt.Println("From Server Msg: ", recv.Message)
}
