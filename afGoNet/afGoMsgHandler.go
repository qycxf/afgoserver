package afGoNet

import (
	"cxfProject/afGo/afGoface"
	"cxfProject/afGo/global"
	"fmt"
	"strconv"
)

type MsgHandle struct {

	//存放每个msgId所对应的处理方法

	Apis map[uint32]afGoface.IRouter

	//负责worker取任务的消息队列
	TaskQueue []chan afGoface.IRequest

	//业务工作worker池的worker数量
	WorkerPoolSize uint32
}

func NewMessageHandle() *MsgHandle {

	return &MsgHandle{
		Apis:           make(map[uint32]afGoface.IRouter),
		WorkerPoolSize: global.Cfg.WorkerPoolSize, //从全局配置中获取
		TaskQueue:      make([]chan afGoface.IRequest, global.Cfg.WorkerPoolSize),
	}
}

//调度执行对应的router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request afGoface.IRequest) {
	//1从request中找到msgId
	handle, ok := mh.Apis[request.GetMsgId()]

	if !ok {
		fmt.Println("api msgId=", request.GetMsgId(), "is Not found!")
	}

	handle.PreHandle(request)
	handle.Handle(request)
	handle.PostHandle(request)

}

//为消息添加具体的逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router afGoface.IRouter) {

	//1.判断 当前的 msg绑定的api处理方法是否已经存在

	if _, ok := mh.Apis[msgId]; ok {
		panic("repeat api,msgId=" + strconv.Itoa(int(msgId)))
		return
	}
	mh.Apis[msgId] = router
	fmt.Println("add api msgId=", msgId, "success")

}

//启动一个worker工作池(只发生一次)
func (mh *MsgHandle) StartWorkerPool() {

	//根据workerSize分别开启worker，每个worker用一个go 承载

	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker启动，当前worker对应的channel 消息队列，开辟空间
		mh.TaskQueue[i] = make(chan afGoface.IRequest, global.Cfg.MaxWorkerSize)
		go mh.StartOneWorker(i, mh.TaskQueue[i])

	}

}

//启动一个worker工作流程
func (mh *MsgHandle) StartOneWorker(workerId int, taskQueue chan afGoface.IRequest) {

	fmt.Println("start worker ！ id:", workerId)

	for {
		select {
		//消息过来，出列，执行当前所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)

		}
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request afGoface.IRequest) {

	//将消息平均分配给不同的worker

	//根据客户端建立的connId来分配
	workerId := request.GetConnection().GetConnID() % mh.WorkerPoolSize

	fmt.Println("connId:", request.GetConnection().GetConnID(), "workerId:", workerId)

	mh.TaskQueue[workerId] <- request

	//将消息发送给对应的taskQueue处理

}
