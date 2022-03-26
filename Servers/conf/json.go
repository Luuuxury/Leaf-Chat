package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel string
	LogPath  string
	// websocket 配置
	WSAddr   string
	CertFile string
	KeyFile  string
	// tcp 配置
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string

	// db conf
	MongodbAddr       string
	MongodbSessionNum int
}

func init() {
	data, err := ioutil.ReadFile("/Users/liuyang/Desktop/Go_Dev/leafserver/bin/conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
