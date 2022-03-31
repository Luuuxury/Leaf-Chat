# Go leaf-server框架从入门到懵逼

## Processor.Register()

在`msg.go`文件里，我们先看

```go
var Processor = json.NewProcessor()
```

json是作者写的一个包，而非官方的`encoding/json` ,   NewProcessor返回一个指向Processor的指针，指针指向了一个`make(map[string]*MsgInfo)`,我的理解大概是先开辟了一个消息模版，后面再通过具体的成员对象进行实例化

```go
type Processor struct {
	msgInfo map[string]*MsgInfo
}

type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}
func NewProcessor() *Processor {
   p := new(Processor)
   p.msgInfo = make(map[string]*MsgInfo)
   return p
}
type MsgHandler func([]interface{})
```

```go
func (p *Processor) Register(msg interface{}) string {
	curr_msgType := reflect.TypeOf(msg)
	...
	msgName := curr_msgType.Elem().Name()
	...
	newMsgInfoStruct := new(MsgInfo)
	newMsgInfoStruct.msgType = curr_msgType
	p.msgInfo[msgName] = newMsgInfoStruct
	return msgName
}
```

我的理解`Register(msg interface{}`，就是将`msg`的消息名作为`map`的Key值，MsgInfo 结构体作为 Val值，返回一个Processor 实例化的对象，这样就记录下了一个固定消息结构体的名称和与其对应的消息内容,也就是我们对应的MsgInfo。

总结：

`Processor.Register(&Hello{})  `就是将结构体名称和  MsgInfo 结构体绑定，而MsgInfo 记录了：

- 该消息结构体的消息类型

- 该消息的路由，路由即该消息下一步由哪个模块处理
- 该消息的业务处理器



<img src="/Users/liuyang/Desktop/Go_Dev/笔记们/leaf_chat.png" style="zoom:150%;" />

## SetRouter()

```go
msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
```

`SetRouter()`逻辑和`Register()`一样，只不过Register()多了个“初始化”，其实都是对`msgInfo`里面成员变量的绑定，`SetRouter()` 接收两个参数，第一个就是我们的消息结构体，第二个可以是任意一个模块的`chanRPC` 

```go
func (p *Processor) SetRouter(msg interface{}, msgRouter *chanrpc.Server) {
	msgType := reflect.TypeOf(msg)
	...
	msgID := msgType.Elem().Name()
	i, ok := p.msgInfo[msgID]
	...
	i.msgRouter = msgRouter
}

// SetRouter 完成了msg.go 中消息结构体和 msgRouter的绑定
type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}
```

总结`SetRouter()`决定了某个消息具体交给内部的哪个模块来处理，这里我们是给了游戏模块进行业务处理。



## Game.Handler()

​	当我们设置完路由的时候，就决定好了这个消息要由那个模块来处理，如果是用户第一次登陆，则就需要注册模块去跟数据库对接来完成用户注册的Handler()，如果这个消息是用户在游戏内发起的聊天信息，则就需要我们用游戏模块来对消息进行 私聊、广播等的Handler()

​	所以，具体由哪个模块处理需要依情况而定，可以是 Login.Handler() 也可以是 Game.Handler()

好的，压力来到game/internal/handler.go这边















## SetHandler() & SetRawHandler()

这两个代码结构也是一样的，同样完成了`msg.go`中`Hello{}`消息结构体和`MsgInfo`中msgHandler和msgRawHandler成员变量的绑定。

最后总结,`Register`其实就是完成了一次消息结构体的解析和对该消息结构体的持久化“注册、绑定“，绑定信息为 `MsgInfo`结构体里的成员变量，而这些成员变量是通过解析传入的`Register(msg interface{})`得到的。

```go
var Processor_json = json.NewProcessor()
func init() {
   Processor_json.Register(&Hello{})
}
```



