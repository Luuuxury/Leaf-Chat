package internal

import (
	"github.com/name5566/leaf/gate"
	conf2 "leaf-chat/Servers/conf"
	"leaf-chat/Servers/game"
	"leaf-chat/Servers/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf2.Server.MaxConnNum,
		PendingWriteNum: conf2.PendingWriteNum,
		MaxMsgLen:       conf2.MaxMsgLen,
		WSAddr:          conf2.Server.WSAddr,
		HTTPTimeout:     conf2.HTTPTimeout,
		CertFile:        conf2.Server.CertFile,
		KeyFile:         conf2.Server.KeyFile,
		TCPAddr:         conf2.Server.TCPAddr,
		LenMsgLen:       conf2.LenMsgLen,
		LittleEndian:    conf2.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
