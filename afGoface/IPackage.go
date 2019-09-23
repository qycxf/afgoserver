package afGoface

type IPackage interface {

	//获取包的头的方法长度
	GetHeadLen() uint32

	//封包方法
	Pack(msg IMessage) ([]byte, error)

	//拆包方法
	Unpack([]byte) (IMessage, error)
}
