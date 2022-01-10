package main

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/knet"
)

type PingRouter struct {
	knet.BaseRouter
}

// PreHandle test prehandle
func (b *PingRouter) PreHandle(request kiface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}

}

// Handle test handle
func (b *PingRouter) Handle(request kiface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("Call back ping...ping...ping error")
	}
}

// PostHandle test posthandle
func (b *PingRouter) PostHandle(request kiface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		fmt.Println("Call router after ping error")
	}
}

func main() {
	s := knet.NewServer("[kinx v0.3]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
