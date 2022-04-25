package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

//var Processor = protobuf.NewProcessor()

type UserRegist struct {
	RegistName string
	RegistPW   string
}

type UserLogin struct {
	LoginName string
	LoginPW   string
}

type RegistResult struct {
	Message string
}

type LoginResult struct {
	Message string
}

// ========== BroadCast Chat ============

func init() {
	Processor.Register(&UserRegist{})
	Processor.Register(&UserLogin{})
	Processor.Register(&RegistResult{})
	Processor.Register(&LoginResult{})

}
