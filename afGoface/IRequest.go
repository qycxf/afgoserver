package afGoface

type IRequest interface {

	//得到链接

	GetConnection() IConnection
	GetData() []byte
}
