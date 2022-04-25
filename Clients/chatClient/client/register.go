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

	registdata := []byte(`{
		"UserRegist": {
			"RegistName": "admin-3",
			"RegistPW": "admin123-3"
		}
	}`)

	buf := make([]byte, 2+len(registdata))
	binary.BigEndian.PutUint16(buf, uint16(len(registdata)))
	copy(buf[2:], registdata)
	_, err = conn.Write(buf)
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
