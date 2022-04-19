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
	//
	RegisterData := []byte(`{
		"UserRegister": {
			"RegisterName": "lyl",
			"RegisterPW": "Fuck"
		}
	}`)
	registerMsg := make([]byte, 2+len(RegisterData))
	//
	binary.BigEndian.PutUint16(registerMsg, uint16((len(RegisterData))))
	copy(registerMsg[2:], RegisterData)
	conn.Write(registerMsg)
}
