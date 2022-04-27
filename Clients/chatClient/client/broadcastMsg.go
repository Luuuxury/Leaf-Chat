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
	broadcastMsg := &msg.C2S_Message{
		Message: "Test Message",
	}

	marshaldata, err := proto.Marshal(broadcastMsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	writeBuf := make([]byte, 2+2+len(marshaldata))
	binary.BigEndian.PutUint16(writeBuf, uint16(2+len(marshaldata)))
	binary.BigEndian.PutUint16(writeBuf[2:], uint16(2))
	copy(writeBuf[4:], marshaldata)
	conn.Write(writeBuf)

}
