package mongodb

import (
	"fmt"
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

// 连接消息
var dialContext = new(mongodb.DialContext)

func init() {
	Connect()
}
func Connect() {
	c, err := mongodb.Dial("localhost", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.EnsureUniqueIndex("userDB", "regist", []string{"name"})
	log.Release("MongoDb 连接成功")
	dialContext = c

	InitData()
}

func InitData() {
	err := Find("userDB", "regist", bson.M{"name": "InitName"})
	if err == nil {
		log.Debug("用户数据库已经初始化过了")
	} else {
		err = Insert("userDB", "regist", bson.M{"name": "InitName", "password": "InitPW"})
		if err != nil {
			log.Debug("用户数据初始化失败: ", err)
		}
	}
}

func Find(db string, collection string, docs interface{}) error {
	c := dialContext
	s := c.Ref()
	defer c.UnRef(s)

	type person struct {
		Id_  bson.ObjectId `bson:"_id"`
		Name string        `bson:"name"`
	}
	user := new(person)
	err := s.DB(db).C(collection).Find(docs).One(&user)
	if err != nil {
		return err
	}
	return err
}

func Insert(db string, collection string, docs interface{}) error {
	c := dialContext
	s := c.Ref()
	defer c.UnRef(s)
	err := s.DB(db).C(collection).Insert(docs)
	if err != nil {
		return err
	}
	return err
}

type UserData struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Password string
}

func FetchUserData(name string) (*UserData, error) {
	c := dialContext
	s := c.Ref()
	defer c.UnRef(s)

	resultData := &UserData{}
	err := s.DB("userDB").C("regist").Find(bson.M{"name": name}).One(resultData)
	return resultData, err
}
