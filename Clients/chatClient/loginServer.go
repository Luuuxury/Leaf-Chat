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

	LoginerData := []byte(`{
	"UserLogin":{
		"LoginName": "lyyyy",
		"LoginPW": "fuck"
		}
	}`)
	loginerMsg := make([]byte, 2+len(LoginerData))
	binary.BigEndian.PutUint16(loginerMsg, uint16((len(LoginerData))))
	copy(loginerMsg[2:], LoginerData)
	conn.Write(loginerMsg)

	//
	//GameData := []byte(`{
	//	"ToGameModuleMsg": {
	//		"MsgInfo": "这个消息需要游戏业务处理"
	//	}
	//}`)
	//// len + data
	//gameMsg := make([]byte, 2+len(GameData))
	//// 默认使用大端序
	//binary.BigEndian.PutUint16(gameMsg, uint16(len(GameData)))
	//copy(gameMsg[2:], GameData)
	//// 发送消息
	//conn.Write(gameMsg)
}
