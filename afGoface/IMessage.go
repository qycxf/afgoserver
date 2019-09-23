package afGoface

type IMessage interface {
	GetMsgId() uint32

	GetMsgLen() uint32

	GetData() []byte

	SetMsgId(id uint32)

	SetMsgLen(len uint32)

	SetData(data []byte)
}
