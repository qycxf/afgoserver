package afGoNet

import (
	"cxfProject/afGo/afGoface"
	"errors"
	"fmt"
	"sync"
)

type ConnectionManage struct {
	connections map[uint32]afGoface.IConnection
	connLock    sync.RWMutex
}

func NewConnManage() *ConnectionManage {

	return &ConnectionManage{

		connections: make(map[uint32]afGoface.IConnection),
	}
}

//添加连接

func (connManage *ConnectionManage) AddConnection(conn afGoface.IConnection) {

	connManage.connLock.Lock()
	defer connManage.connLock.Unlock()
	len := connManage.GetConnectionLen()
	connManage.connections[conn.GetConnID()] = conn
	fmt.Println("add conn success! add num:", connManage.GetConnectionLen()-len)
}

//删除连接
func (connManage *ConnectionManage) RemoveConnection(conn afGoface.IConnection) {

	connManage.connLock.Lock()
	defer connManage.connLock.Unlock()

	len := connManage.GetConnectionLen()

	delete(connManage.connections, conn.GetConnID())

	fmt.Println("remove conn success! remove num:", len-connManage.GetConnectionLen())
}

//根据connId获取连接
func (connManage *ConnectionManage) GetConnection(connId uint32) (afGoface.IConnection, error) {

	connManage.connLock.RLock()
	defer connManage.connLock.RUnlock()

	if conn, ok := connManage.connections[connId]; ok {
		return conn, nil
	}

	return nil, errors.New("not found conn")
}

//获取 连接数
func (connManage *ConnectionManage) GetConnectionLen() int {

	return len(connManage.connections)
}

//清除并终止所有的连接
func (connManage *ConnectionManage) ClearConnection() {
	connManage.connLock.Lock()
	defer connManage.connLock.Unlock()

	for connId, conn := range connManage.connections {
		conn.Stop()
		delete(connManage.connections, connId)
	}

	fmt.Println("clear all connections success!")

}
