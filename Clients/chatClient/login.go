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

	logindata := []byte(`{
	"UserLogin":{
		"LoginName": "admin",
		"LoginPW": "admin123"
		}
	}`)
	buf := make([]byte, 2+len(logindata))
	binary.BigEndian.PutUint16(buf, uint16(len(logindata)))
	copy(buf[2:], logindata)
	conn.Write(buf)
}
