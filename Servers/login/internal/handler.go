package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"leaf-chat/Servers/db/mongodb"
	"leaf-chat/Servers/msg"
	"reflect"
	"regexp"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//handler(&msg.UserRegister{}, handleLogin)
	handler(&msg.UserLogin{}, handleRegister)
}

func handleRegister(args []interface{}) {
	receMsg := args[0].(*msg.UserRegister)
	agent := args[1].(gate.Agent)
	returnMsg := &msg.UserRegisterResult{}
	log.Debug("receive UserRegister name=%v", receMsg.RegisterName)

	// UserName 注册规则
	reg := regexp.MustCompile(`/^[a-zA-Z\d]\w{2,10}[a-zA-Z\d]$/`)
	matched := reg.FindString(receMsg.RegisterName)
	if matched != " " {
		log.Debug("注册用户名不合法")
	}
	// 判断用户是否已经注册
	err := mongodb.Find("game", "login", bson.M{"name": receMsg.RegisterName})
	if err == nil {
		fmt.Println(err)
		log.Debug("该用户名已经注册, 注册失败")
		returnMsg.Err = "用户名已经注册了, 请重新注册～"
		returnMsg.Retresult = "Retresult is ok"
		agent.WriteMsg(returnMsg)
	}
}

/*
// 消息处理
func handleLogin(args []interface{}) {
	receMsg := args[0].(*msg.UserLogin)
	agent := args[1].(gate.Agent)

	sendMsg := &msg.UserLoginResult{}
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
*/
