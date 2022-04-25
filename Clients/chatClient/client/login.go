package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"leaf-chat/Servers/msg"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:21002")
	if err != nil {
		log.Debug("客户端连接失败: ", err)
	}

	clientdata := &msg.UserLogin{
		LoginName: "admin-1",
		LoginPW:   "admin-1",
	}
	protoMarshal, err := proto.Marshal(clientdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	//  -------------------------------
	//  | len | id | protobuf message |
	//  -------------------------------
	writeBuf := make([]byte, 2+2+len(protoMarshal)) // 2: 记录数据长度 ， 2：id-2字节

	// 默认使用大端序
	binary.BigEndian.PutUint16(writeBuf, uint16(2+len(writeBuf))) // [0:2] 俩字节记录 整个protobuf长度， 长度包括 len(id) + len(data)
	binary.BigEndian.PutUint16(writeBuf[2:], uint16(0))           // [2:4] 俩字节记录 id 编号是多少
	copy(writeBuf[4:], protoMarshal)
	// 发送消息
	conn.Write(writeBuf)

	for {
		readBuf := make([]byte, 4096)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println("与服务器断开连接")
			break
		}
		registResult := string(readBuf[:n])
		fmt.Println(registResult)
	}

}
