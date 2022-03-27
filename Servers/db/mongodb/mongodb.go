package mongodb

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"leaf-chat/Servers/conf"
)

var (
	Context *mongodb.DialContext
)

func init() {
	var err error
	Context, err = mongodb.Dial(conf.Server.MongodbAddr, conf.Server.MongodbSessionNum)
	if err != nil {
		log.Fatal("mongondb init is error(%v)", err)
	}
}
