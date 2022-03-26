package game

import (
	"leaf-chat/Servers/game/internal"
)

var (
	// 实例化 game 模块
	Module = new(internal.Module)
	// 暴露 ChanRPC
	ChanRPC = internal.ChanRPC
)
