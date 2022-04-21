package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

type UserRegist struct {
	RegistName string
	RegistPW   string
}

type UserLogin struct {
	LoginName string
	LoginPW   string
}

// ========== BroadCast Chat ============

type C2S_Message struct {
	Message string
}

type S2C_Message struct {
	UserName string
	Message  string
}

func init() {
	Processor.Register(&UserRegist{})
	Processor.Register(&UserLogin{})

	Processor.Register(&C2S_Message{})
	Processor.Register(&S2C_Message{})

}
