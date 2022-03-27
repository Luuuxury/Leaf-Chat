package gate

import (
	"leaf-chat/Servers/game"
	"leaf-chat/Servers/login"
	"leaf-chat/Servers/msg"
)

func init() {
	// 登录消息(Login)由login模块处理，此处类似于django的urls.py
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.ToGameModuleMsg{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.UserLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.UserRegister{}, login.ChanRPC)
}

// 一切就绪，我们现在可以在 game 模块中处理 Hello 消息了。打开 LeafServer game/internal/handler.go，敲入如下代码：
