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

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.UserRegist{}, handleRegist)
	handler(&msg.C2S_Message{}, handleBroadCast)
	handler(&msg.UserLogin{}, handleLogin)
}

var loc = time.FixedZone("", 8*3600)

var onlineMap = make(map[gate.Agent]struct{})

func broadcast(msg interface{}, _a gate.Agent) {
	for a := range onlineMap {
		if a == _a {
			continue
		}
		a.WriteMsg(msg)
	}
}

func handleRegist(args []interface{}) {
	recv := args[0].(*msg.UserRegist)
	agent := args[1].(gate.Agent)
	//判断用户是否已经注册
	err := mongodb.Find("userDB", "regist", bson.M{"name": recv.RegistName})
	if err == nil {
		retResult := &msg.RegistResult{
			Message: "该用户名已经被注册",
		}
		agent.WriteMsg(retResult)
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
		//log.Debug("数据库添加用户名失败!")
		retResult := &msg.RegistResult{
			Message: "注册失败，请重新注册",
		}
		agent.WriteMsg(retResult)
		return
	} else {
		//log.Debug("数据库添加用户成功!")
		retResult := &msg.RegistResult{
			Message: "注册成功",
		}
		agent.WriteMsg(retResult)
		return
	}
}

func handleLogin(args []interface{}) {
	recv := args[0].(*msg.UserLogin)
	agent := args[1].(gate.Agent)
	onlineMap[agent] = struct{}{}
	agent.SetUserData(recv.LoginName)
	if recv.LoginName == "" {
		//log.Debug("数据库添加用户成功!")
		retResult := &msg.LoginResult{
			Message: "请输入正确的用户名",
		}
		agent.WriteMsg(retResult)
		return
	}
	// 数据库获取该用户信息
	userData, err := mongodb.FetchUserData(recv.LoginName)
	if err == mgo.ErrNotFound {
		//log.Debug("登陆用户名不存在")
		retResult := &msg.LoginResult{
			Message: "登陆用户名不存在",
		}
		agent.WriteMsg(retResult)
		return
	}

	// 密码核对
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(recv.LoginPW))
	if err != nil {
		//log.Debug("输入的用户名或密码不正确!")
		retResult := &msg.LoginResult{
			Message: "输入的用户名或密码不正确",
		}
		agent.WriteMsg(retResult)
		return
	} else {
		retResult := &msg.LoginResult{
			Message: "登陆成功",
		}
		agent.WriteMsg(retResult)
		// 广播用户上线
		time.Sleep(time.Second * 1)
		broacastMsg := &msg.S2C_Message{
			UserName: recv.LoginName,
			Message:  "用户上线",
		}
		broadcast(broacastMsg, agent)
	}
}

func handleBroadCast(args []interface{}) {
	recv := args[0].(*msg.C2S_Message)
	agent := args[1].(gate.Agent)
	now := time.Now().In(loc)
	message := fmt.Sprintf("@%02d:%02d %s", now.Hour(), now.Minute(), recv.Message)

	broacastMsg := &msg.S2C_Message{
		UserName: "消息中心",
		Message:  message,
	}
	broadcast(broacastMsg, agent)

}
