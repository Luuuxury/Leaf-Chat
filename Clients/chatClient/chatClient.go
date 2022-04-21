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
	"C2S_AddUser":{
		"UserName": "lyl"
		}
	}`)

	loginerMsg := make([]byte, 2+len(LoginerData))
	binary.BigEndian.PutUint16(loginerMsg, uint16(len(LoginerData)))
	copy(loginerMsg[2:], LoginerData)
	conn.Write(loginerMsg)
}
