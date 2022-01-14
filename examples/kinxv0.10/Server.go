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

func DoConnBegin(conn kiface.IConnection) {
	fmt.Println("===>DoConnBegin is Called...")
	if err := conn.SendMsg(202, []byte("DoConn Begin")); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Set conn Name...")
	conn.SetProperty("Name", "testName")

}

func DoConnAfter(conn kiface.IConnection) {
	fmt.Println("===>DoConnAfter is Called...")
	fmt.Println("conn ID = ", conn.GetConnID(), "is lost...")
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}
}

func main() {
	s := knet.NewServer()
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	s.SetOnConnStart(DoConnBegin)
	s.SetOnConnStop(DoConnAfter)

	s.Serve()
}
