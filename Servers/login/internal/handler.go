package internal

import (
	"github.com/name5566/leaf/gate"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"leaf-chat/Servers/db/mongodb/accountDB"
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
	// 获取该人员的数据库信息
	accountData, err := accountDB.Get(receMsg.LoginName)

	if err == mgo.ErrNotFound {
		accountData = &accountDB.Data{Id: bson.NewObjectId(), Name: receMsg.LoginName, Password: receMsg.LoginPW}
		err = accountDB.Create(accountData)
	}
	if err != nil {
		sendErrFunc(err.Error())
		return
	} else if accountData.Password != receMsg.LoginPW {
		sendErrFunc("password is error")
		return
	}

}
