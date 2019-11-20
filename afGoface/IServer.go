package afGoface

type IServerFace interface {
	Start()
	Stop()
	Server()

	GetConnMange() IConnectionManage

	//路由功能，给当前的服务注册一个路由功能，给客户端的链接使用
	AddRouter(msgId uint32, router IRouter)

	//注册hookStart方法

	SetOnConnStart(func(conn IConnection))
	//注册hookStop方法
	SetOnConnStop(func(conn IConnection))
	//调用hookStart方法
	CallOnConnStart(conn IConnection)
	//调用hookStop方法
	CallOnConnStop(conn IConnection)
}
