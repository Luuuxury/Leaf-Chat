package internal

import (
	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var users = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	users[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(users, a)

	_, ok := a.UserData().(string)
	if !ok {
		return
	}
}

func broadcastMsg(msg interface{}, _a gate.Agent) {
	for a := range users {
		if a == _a {
			continue
		}
		a.WriteMsg(msg)
	}
}
