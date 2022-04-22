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

	//判断用户是否已经注册
	err := mongodb.Find("userDB", "regist", bson.M{"name": recv.RegistName})
	if err == nil {
		retResult := &msg.RegistResult{
			Message: "该用户名已经被注册",
		}
		agent.WriteMsg(retResult)
	}
	// 如果该用户名没有被注册过，就直接 insert

	hashPw, err := bcrypt.GenerateFromPassword([]byte(recv.RegistPW), bcrypt.DefaultCost)
	if err != nil {
		log.Debug("密码加密过程出错")
	}
	strPw := string(hashPw)
	err = mongodb.Insert("userDB", "regist", bson.M{"name": recv.RegistName, "password": strPw})
	if err != nil {
		//log.Debug("数据库添加用户名失败!")
		retResult := &msg.RegistResult{
			Message: "注册失败，请重新注册",
		}
		agent.WriteMsg(retResult)
	} else {
		//log.Debug("数据库添加用户成功!")
		retResult := &msg.RegistResult{
			Message: "注册成功",
		}
		agent.WriteMsg(retResult)
	}
}

func handleLogin(args []interface{}) {
	recv := args[0].(*msg.UserLogin)
	agent := args[1].(gate.Agent)

	if recv.LoginName == "" {
		//log.Debug("数据库添加用户成功!")
		retResult := &msg.RegistResult{
			Message: "请输入正确的用户名",
		}
		agent.WriteMsg(retResult)
		return
	}

	// 数据库获取该用户信息
	userData, err := mongodb.FetchUserData(recv.LoginName)
	if err == mgo.ErrNotFound {
		//log.Debug("登陆用户名不存在")
		retResult := &msg.RegistResult{
			Message: "登陆用户名不存在",
		}
		agent.WriteMsg(retResult)
	}
	// 密码核对
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(recv.LoginPW))
	if err != nil {
		//log.Debug("输入的用户名或密码不正确!")
		retResult := &msg.RegistResult{
			Message: "输入的用户名或密码不正确",
		}
		agent.WriteMsg(retResult)
	} else {
		retResult := &msg.RegistResult{
			Message: "登陆成功",
		}
		agent.WriteMsg(retResult)

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
