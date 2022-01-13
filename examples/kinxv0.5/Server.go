package main

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/knet"
)

type PingRouter struct {
	knet.BaseRouter
}

// Handle test handle
func (b *PingRouter) Handle(request kiface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client: msgID= ", request.GetMsgID(),
		", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping ping ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := knet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
