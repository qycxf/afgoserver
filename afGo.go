package main

import (
	"cxfProject/afGo/afGoNet"
	"cxfProject/afGo/afGoface"
	"fmt"
)

type PingRouter struct {
	afGoNet.BaseRouter
	data []byte
}

func (this *PingRouter) PreHandle(request afGoface.IRequest) {

	fmt.Println("Call before Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))

	if err != nil {
		fmt.Println("Call Router PreHandle err")
	}

}

func (this *PingRouter) Handle(request afGoface.IRequest) {

	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping \n"))

	if err != nil {
		fmt.Println("Call Router Handle err")
	}

}
func (this *PingRouter) PostHandle(request afGoface.IRequest) {
	fmt.Println("Call after Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n "))

	if err != nil {
		fmt.Println("Call Router PostHandle err")
	}
}

type HelloRouter struct {
	afGoNet.BaseRouter
	data []byte
}

func (this *HelloRouter) PreHandle(request afGoface.IRequest) {

}

func (this *HelloRouter) Handle(request afGoface.IRequest) {

	fmt.Println("Call Router Handle...")
	err := request.GetConnection().SendMsg(201, []byte("hello afGo1.0.6"))

	if err != nil {
		fmt.Println("Call Router Handle err")
	}

}
func (this *HelloRouter) PostHandle(request afGoface.IRequest) {

}
func main() {

	s := afGoNet.NewServer("afGo[1.0.5]")

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Server()
}
