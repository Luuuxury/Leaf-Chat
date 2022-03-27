package accountDB

import (
	lmongodb "github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"leaf-chat/Servers/db/mongodb"
)

type Data struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Password string
}

func init() {
	session := mongodb.Context.Ref()
	defer mongodb.Context.UnRef(session)

	getCollection(session).EnsureIndex(mgo.Index{
		Key:    []string{"name"},
		Unique: true,
		Sparse: true,
	})
}

func getCollection(session *lmongodb.Session) *mgo.Collection {
	// DB("数据库名") C - collection ("表名/ 集合名")
	return session.DB("login").C("account")
}

// Get 从db中获取这个人
func Get(name string) (*Data, error) {
	session := mongodb.Context.Ref()
	defer mongodb.Context.UnRef(session)

	result := &Data{}
	err := getCollection(session).Find(bson.M{"name": name}).One(result)
	return result, err
}

// Create 数据库创建这个人
func Create(account *Data) error {
	//  Id bson.ObjectId `bson:"_id"`
	//	Name     string
	//	Password string
	session := mongodb.Context.Ref()
	defer mongodb.Context.UnRef(session)

	return getCollection(session).Insert(account)
}
