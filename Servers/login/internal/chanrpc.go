package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var onlineMap = make(map[gate.Agent]struct{})

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	onlineMap[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(onlineMap, a)

}

func broadcastMsg(msg interface{}, _a gate.Agent) {
	for a := range onlineMap {
		if a == _a {
			log.Debug("群发不用再给自己发了")
			continue
		}
		a.WriteMsg(msg)
	}
}
