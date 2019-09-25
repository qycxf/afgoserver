package afGoNet

import (
	"cxfProject/afGo/afGoface"
	"fmt"
	"strconv"
)

type MsgHandle struct {

	//存放每个msgId所对应的处理方法

	Apis map[uint32]afGoface.IRouter
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
