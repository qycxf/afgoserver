package afGoface

type IRouter interface {

	//处理业务之前

	PreHandle(request IRequest)

	//在处理conn业务的主方法
	Handle(reauest IRequest)

	//在处理conn之后的方法
	PostHandle(request IRequest)
}
