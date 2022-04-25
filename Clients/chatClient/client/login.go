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
	logindata := &msg.UserLogin{
		LoginName: "admin-1",
		LoginPW:   "admin-1",
	}
	marshaldata, err := proto.Marshal(logindata)
	if err != nil {
		fmt.Println(err)
		return
	}

	writeBuf := make([]byte, 2+2+len(marshaldata))
	binary.BigEndian.PutUint16(writeBuf, uint16(2+len(marshaldata)))
	binary.BigEndian.PutUint16(writeBuf[2:], uint16(1))
	copy(writeBuf[4:], marshaldata)
	// 发送消息
	conn.Write(writeBuf)

	// 接收服务端消息
	for {
		// 接收消息
		readBuf := make([]byte, 1024)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println("读取服务端业务处理结果失败!")
			break
		}
		recv := &msg.LoginResult{}
		err = proto.Unmarshal(readBuf[4:n], recv)
		if err != nil {
			log.Debug("unmarshaling error: ", err)
		}
		fmt.Println(recv.Message)
	}

}
