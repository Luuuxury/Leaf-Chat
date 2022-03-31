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
	//defer c.Close()
	c.EnsureUniqueIndex("game", "login", []string{"name"})
	log.Release("mongodb Connect success")
	dialContext = c

	InitData()
}

func InitData() {
	err := Find("game", "login", bson.M{"name": "InitName"})
	if err == nil {
		log.Debug("数据库已经初始化过了", err)
	} else {
		err = Insert("game", "login", bson.M{"name": "InitName", "password": "qq123456"})
		if err != nil {
			fmt.Println(err)
			log.Debug("数据初始化插入失败", err)
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
		fmt.Println(err)
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
		fmt.Println(err)
		return err
	}
	return err
}
