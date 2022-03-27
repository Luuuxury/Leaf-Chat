package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

type ToGameModuleMsg struct {
	Name string
}

type UserRegister struct {
	RegisterName string
	RegisterPW   string
}

type UserRegisterResult struct {
	Err       string
	Retresult string
}

type UserLogin struct {
	LoginName string
	LoginPW   string
}

type UserLoginResult struct {
	Err       string
	Retresult string
}

func init() {
	Processor.Register(&ToGameModuleMsg{})
	Processor.Register(&UserLogin{})
	Processor.Register(&UserRegister{})
}
