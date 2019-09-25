package afGoface

type IServerFace interface {
	Start()
	Stop()
	Server()

	//路由功能，给当前的服务注册一个路由功能，给客户端的链接使用
	AddRouter(msgId uint32, router IRouter)
}
