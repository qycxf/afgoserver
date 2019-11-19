package afGoface

type IMsgHandle interface {

	//调度执行对应的router消息处理方法
	DoMsgHandler(request IRequest)

	//为消息添加具体的逻辑
	AddRouter(msgId uint32, router IRouter)

	//启动worker工作池
	StartWorkerPool()

	//将请求交给消息处理队列
	SendMsgToTaskQueue(request IRequest)
}
