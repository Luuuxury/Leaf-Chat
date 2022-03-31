package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"leaf-chat/Servers/db/mongodb"
	"leaf-chat/Servers/msg"
	"reflect"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//handler(&msg.UserRegister{}, handleLogin)
	handler(&msg.UserRegister{}, handleRegister)
}

func handleRegister(args []interface{}) {
	receMsg := args[0].(*msg.UserRegister)
	agent := args[1].(gate.Agent)

	returnMsg := &msg.UserRegisterResult{}
	log.Debug("receive RegisterName is %v", receMsg.RegisterName)
	log.Debug("receive RegisterPW is %v", receMsg.RegisterPW)

	//判断用户是否已经注册
	err := mongodb.Find("game", "login", bson.M{"name": receMsg.RegisterName})
	if err == nil {
		fmt.Println("执行 mongodb.Find 完成， err为None, err is", err)
		log.Debug("Debug is 该用户名已经注册, 请换个用户名")
		returnMsg.Err = "用户名已经注册了, 请重新注册～"
		returnMsg.Retresult = "Retresult is ok"
		// 给客户端返回，说明已经该用户已经注册过了
		agent.WriteMsg(returnMsg)
	}
	// 如果该用户名没有被注册过，就直接 insert
	err = mongodb.Insert("game", "login", bson.M{"name": receMsg.RegisterName, "password": receMsg.RegisterPW})
	if err != nil {
		fmt.Println("执行插入用户操作失败, 报错提示 is", err)
		log.Debug("Debug 用户名写入失败!")
		returnMsg.Err = "returnMsg.Err is :用户名插入失败！"
		returnMsg.Retresult = "Retresult is ok"
		agent.WriteMsg(returnMsg)
	} else {
		fmt.Println("执行插入用户操作 sucess")
		log.Debug("Debug UserRegister write in success")
		returnMsg.Err = "returnMsg.Err is :用户名插入成功！"
		returnMsg.Retresult = "Retresult is ok"
		agent.WriteMsg(returnMsg)
	}
}

/*
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
