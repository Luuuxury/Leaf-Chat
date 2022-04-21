package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"leaf-chat/Servers/db/mongodb"
	"leaf-chat/Servers/msg"
	"reflect"
	"time"
)

const maxMessages = 50

var (
	messages [maxMessages]struct {
		userName string
		message  string
	}
	messageIndex int
)

var loc = time.FixedZone("", 8*3600)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.UserRegist{}, handleRegist)
	handler(&msg.UserLogin{}, handleLogin)

	handler(&msg.C2S_Message{}, handleC2SMessage)
}

func handleRegist(args []interface{}) {
	recv := args[0].(*msg.UserRegist)
	//agent := args[1].(gate.Agent)
	log.Debug("注册用户名是: %v", recv.RegistName)
	log.Debug("注册密码是: %v", recv.RegistPW)

	//判断用户是否已经注册
	err := mongodb.Find("userDB", "regist", bson.M{"name": recv.RegistName})
	if err == nil {
		log.Debug("该用户名已经被注册, 请换个用户名!")
		return
	}
	// 如果该用户名没有被注册过，就直接 insert
	err = mongodb.Insert("userDB", "regist", bson.M{"name": recv.RegistName, "password": recv.RegistPW})
	if err != nil {
		log.Debug("数据库添加用户名失败!")
	} else {
		log.Debug("数据库添加用户成功!")
	}
}

func handleLogin(args []interface{}) {
	receMsg := args[0].(*msg.UserLogin)
	agent := args[1].(gate.Agent)
	log.Debug("登陆用户名是: %v", receMsg.LoginName)
	log.Debug("登陆密码是: %v", receMsg.LoginPW)
	if receMsg.LoginName == "" {
		log.Debug("登陆用户名为空!")
		return
	}

	// 从数据库核对用户名和密码
	userData, err := mongodb.FetchUserData(receMsg.LoginName)
	// 如果数据库没有这个人，就把把这个人添加进去
	if err == mgo.ErrNotFound {
		log.Debug("登陆用户名不存在，请输入正确的用户名")
	}
	if err != nil {
		log.Debug("登陆获取用户信息出错 err: %v", err)
		return
		// 如果有这个人，但是密码错误
	} else if userData.Password != receMsg.LoginPW {
		log.Debug("登陆密码不对！")
	}

	// 将该用户加入群聊
	agent.SetUserData(receMsg.LoginName)
	for i := 0; i < maxMessages; i++ {
		index := (messageIndex + i) % maxMessages
		pm := &messages[index]
		if pm.message == "" {
			continue
		}
		agent.WriteMsg(&msg.S2C_Message{
			UserName: pm.userName,
			Message:  pm.message,
		})
	}
}

func handleC2SMessage(args []interface{}) {
	recv := args[0].(*msg.C2S_Message)
	agent := args[1].(gate.Agent)
	userName, ok := agent.UserData().(string)
	if !ok {
		return
	}

	now := time.Now().In(loc)
	message := fmt.Sprintf("@%02d:%02d %s", now.Hour(), now.Minute(), recv.Message)
	pm := &messages[messageIndex]
	pm.userName = userName
	pm.message = message
	messageIndex++

	if messageIndex >= maxMessages {
		messageIndex = 0
	}

	broadcastMsg(&msg.S2C_Message{
		UserName: userName,
		Message:  message,
	}, agent)

}
