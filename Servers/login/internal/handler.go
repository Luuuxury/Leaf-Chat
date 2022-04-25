package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"leaf-chat/Servers/db/mongodb"
	"leaf-chat/Servers/msg"
	"reflect"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.UserRegist{}, handleRegist)
	handler(&msg.UserLogin{}, handleLogin)

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
	//agent.SetUserData(recv.LoginName)

	// 用户上线广播
	broadcastMsg(&msg.LoginResult{
		Message: "新用户上线了",
	}, agent)

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
	}

}
