package afGoface

//连接管理模块

type IConnectionManage interface {

	//添加连接

	AddConnection(conn IConnection)
	//删除连接
	RemoveConnection(conn IConnection)
	//根据connId获取连接
	GetConnection(connId uint32) (IConnection, error)
	//获取 连接数
	GetConnectionLen() int
	//清除并终止所有的连接
	ClearConnection()
}
