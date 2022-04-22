package main

import (
	"encoding/binary"
	"fmt"
	"github.com/name5566/leaf/log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:21002")
	if err != nil {
		log.Debug("客户端连接失败: ", err)
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
	_, err = conn.Write(writeBuf)
	if err != nil {
		fmt.Println("客户端写入数据出错了")
	}

	// ==================== 接收Server消息 ====================
	// 1. 用户接收用户登陆流程信息
	// 2. 登陆后开始接收世界消息

	for {
		readBuf := make([]byte, 4096)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println("服务器下线了")
			break
		}
		registResult := string(readBuf[:n])
		fmt.Println(registResult)
	}

}
