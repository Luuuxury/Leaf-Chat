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
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.ToGameModuleMsg{}, handleHello)
}

func handleHello(args []interface{}) {
	receMsg := args[0].(*msg.ToGameModuleMsg)
	// 消息的发送者
	agent := args[1].(gate.Agent)
	// 输出收到的消息的内容
	log.Debug("游戏业务处理模块 接收到了 来自前端的消息 %v", receMsg.MsgInfo)
	agent.WriteMsg(&msg.ToGameModuleMsg{
		MsgInfo: "WriteMsg to Client Call back from game handler!",
	})
}

// 到这里，一个简单的范例就完成了。为了更加清楚的了解消息的格式，我们从 0 编写一个最简单的测试客户端。
