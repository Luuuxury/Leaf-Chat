package gate

import (
	"leaf-chat/Servers/login"
	"leaf-chat/Servers/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.UserRegist{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.UserLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Message{}, login.ChanRPC)
}
