package afGoNet

import "afGo/afGoface"

//实现router时，先嵌入baseRouter基类，然后根据需求对这个基类进行重写就好了
type BaseRouter struct {
}

//这里之所以baseRouter的方法都为空，是因为有的router不希望有preHandle和postHandle，
//所以Router全部继承baseRouter的好处就是  不需要实现preHandle和postHandle，

func (br *BaseRouter) PreHandle(request afGoface.IRequest) {

}

func (br *BaseRouter) Handle(request afGoface.IRequest) {

}

func (br *BaseRouter) PostHandle(request afGoface.IRequest) {

}
