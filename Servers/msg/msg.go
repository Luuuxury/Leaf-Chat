package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&UserRegist{})  // 0
	Processor.Register(&UserLogin{})   // 1
	Processor.Register(&C2S_Message{}) // 2

	Processor.Register(&RegistResult{})
	Processor.Register(&LoginResult{})
	Processor.Register(&S2C_Message{})
}
