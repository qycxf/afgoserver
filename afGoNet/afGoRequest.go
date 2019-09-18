package afGoNet

import "afGo/afGoface"

type Request struct {
	conn afGoface.IConnection

	data []byte
}

func (r *Request) GetConnection() afGoface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {

	return r.data
}
