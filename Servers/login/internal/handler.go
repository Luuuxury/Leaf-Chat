package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"golang.org/x/crypto/bcrypt"
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
	// 收到的消息
	recv := args[0].(*msg.UserRegist)
	// 消息的发送者
	agent := args[1].(gate.Agent)
	log.Debug("注册用户名是: %v", recv.RegistName)
	log.Debug("注册密码是: %v", recv.RegistPW)

	//判断用户是否已经注册
	err := mongodb.Find("userDB", "regist", bson.M{"name": recv.RegistName})
	if err == nil {
		log.Debug("该用户名已经被注册, 请换个用户名!")
		return
	}
	// 如果该用户名没有被注册过，就直接 insert
	hashPw, err := bcrypt.GenerateFromPassword([]byte(recv.RegistPW), bcrypt.DefaultCost)
	if err != nil {
		log.Debug("密码加密过程出错")
	}
	strPw := string(hashPw)
	err = mongodb.Insert("userDB", "regist", bson.M{"name": recv.RegistName, "password": strPw})
	if err != nil {
		log.Debug("数据库添加用户名失败!")
	} else {
		log.Debug("数据库添加用户成功!")
	}

	retBuf := &msg.RegistResult{
		Message: "服务端业务处理完成结果",
	}

	agent.WriteMsg(retBuf)
}

func handleLogin(args []interface{}) {
	recv := args[0].(*msg.UserLogin)
	agent := args[1].(gate.Agent)
	log.Debug("登陆用户名是: %v", recv.LoginName)
	log.Debug("登陆密码是: %v", recv.LoginPW)
	if recv.LoginName == "" {
		log.Debug("登陆用户名为空!")
		return
	}

	// 数据库获取该用户信息
	userData, err := mongodb.FetchUserData(recv.LoginName)
	if err == mgo.ErrNotFound {
		log.Debug("登陆用户名不存在，请输入正确的用户名!")
		return
	}
	// 密码核对
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(recv.LoginPW))
	if err != nil {
		log.Debug("输入的用户名或密码不正确!")
	} else {
		log.Debug("登陆成功!")
	}

	// 将该用户添加到世界聊天
	agent.SetUserData(recv.LoginName)
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
