package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"leaf-chat/Servers/conf"
	"leaf-chat/Servers/game"
	"leaf-chat/Servers/gate"
	"leaf-chat/Servers/login"
)

func main() {
	// 服务器配置文件初始化
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	// 启动各个模块
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
