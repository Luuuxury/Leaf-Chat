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

	registdata := &msg.UserRegist{
		RegistName: "admin-1",
		RegistPW:   "admin-1",
	}
	marshaldata, err := proto.Marshal(registdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	writeBuf := make([]byte, 2+2+len(marshaldata))
	binary.BigEndian.PutUint16(writeBuf[0:2], uint16(2+len(marshaldata)))
	binary.BigEndian.PutUint16(writeBuf[2:4], uint16(0))
	copy(writeBuf[4:], marshaldata)
	conn.Write(writeBuf)

	// 读数据
	for {
		// 接收消息
		readBuf := make([]byte, 1024)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println("读取服务端业务处理结果失败!")
			break
		}
		recv := &msg.RegistResult{}
		err = proto.Unmarshal(readBuf[4:n], recv)
		if err != nil {
			log.Debug("unmarshaling error: ", err)
		}
		fmt.Println(recv.Message)
	}
}
