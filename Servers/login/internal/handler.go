package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"leaf-chat/Servers/msg"
	"reflect"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.C2L_Login{}, handleLogin)
}

// 消息处理
func handleLogin(args []interface{}) {
	receMsg := args[0].(*msg.C2L_Login)
	agent := args[1].(gate.Agent)

	sendMsg := &msg.L2C_Login{}
	sendErrFunc := func(err string) {
		sendMsg.Err = err
		agent.WriteMsg(sendMsg)
	}

	if receMsg.LoginName == "" {
		sendErrFunc("account name is null")
		return
	}
	////后台console 输出内容
	log.Debug("User Login Name is  %v", receMsg.LoginName)
	log.Debug("User Login PW is  %v", receMsg.LoginPW)
	//// 给发送者回应一个 Test 消息
	//a.WriteMsg(&msg.UserLogin{
	//	LoginName: "from fucn handleLogin ",
	//})

	agent.WriteMsg(agent)
}
