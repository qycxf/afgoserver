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
func main() {

	s := afGoNet.NewServer("afGo[1.0.2]")

	s.AddRouter(&PingRouter{})
	s.Server()
}
