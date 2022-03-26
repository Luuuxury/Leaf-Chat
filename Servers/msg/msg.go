package msg

import (
	"github.com/name5566/leaf/network/json"
	"gopkg.in/mgo.v2/bson"
)

var Processor = json.NewProcessor()

type Hello struct {
	Name string
}

type C2L_Login struct {
	LoginName string
	LoginPW   string
}

type L2C_Login struct {
	Err       string
	Id        bson.ObjectId
	FrontAddr string
	Token     bson.ObjectId
}

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&C2L_Login{})
	Processor.Register(&L2C_Login{})
}
