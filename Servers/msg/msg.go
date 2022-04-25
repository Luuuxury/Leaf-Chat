package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&UserRegist{})
	Processor.Register(&UserLogin{})
	Processor.Register(&RegistResult{})
	Processor.Register(&LoginResult{})
}
