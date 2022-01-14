package main

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/knet"
)

type PingRouter struct {
	knet.BaseRouter
}

type HelloRouter struct {
	knet.BaseRouter
}

// Handle test handle
func (b *PingRouter) Handle(request kiface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client: msgID= ", request.GetMsgID(),
		", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("ping ping ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

// Handle test handle
func (b *HelloRouter) Handle(request kiface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client: msgID= ", request.GetMsgID(),
		", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("hello, welcome to kinx!"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := knet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
