package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"leaf-chat/Servers/msg"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:21002")
	if err != nil {
		log.Debug("客户端连接失败: ", err)
	}
	logindata := &msg.UserLogin{
		LoginName: "admin-3",
		LoginPW:   "admin-3",
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

	// 接收登陆消息
	time.Sleep(time.Second * 1)
	readBuf := make([]byte, 1024)
	n, err := conn.Read(readBuf)
	if err != nil {
		fmt.Println("与服务器断开连接")
	}
	recv := &msg.LoginResult{}
	err = proto.Unmarshal(readBuf[4:n], recv)
	if err != nil {
		log.Debug("Client接收消息反序列化出错: ", err)
	}
	fmt.Println(recv.Message)

	// 接收大厅广播消息消息
	for {
		// 接收消息
		readBuf := make([]byte, 1024)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println("与服务器断开连接")
			break
		}
		recv := &msg.S2C_Message{}
		err = proto.Unmarshal(readBuf[4:n], recv)
		if err != nil {
			log.Debug("Client接收消息反序列化出错: ", err)
		}
		fmt.Println(recv.UserName, recv.Message)
	}
}
