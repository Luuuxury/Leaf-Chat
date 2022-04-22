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
		"LoginName": "admin-1",
		"LoginPW": "admin123-1"
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

	for {
		readBuf := make([]byte, 4096)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println("读取服务端业务处理结果失败!")
			break
		}
		registResult := string(readBuf[:n])
		fmt.Println(registResult)
	}

}
