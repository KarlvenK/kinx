package knet

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"net"
)

//Server impl IServer interface
//defines a Server module
type Server struct {
	//server name
	Name string
	//binding ip version
	IPVersion string
	//listened IP
	IP string
	//listened port
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP:%s, Port %d, is starting\n", s.IP, s.Port)
	go func() {
		//get TCP addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		//listen server addr
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listener ", s.IPVersion, " err", err)
			return
		}
		fmt.Println("start kinx server ", s.Name, " succ, Listening...")

		//block client connection handle client's work
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}
					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					if _, err := conn.Write(buf[0:cnt]); err != nil {
						fmt.Println("write back buf err ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	//todo
}

func (s *Server) Serve() {
	//start server services
	s.Start()

	//todo

	//block
	select {}

}

// NewServer init Server module
func NewServer(name string) kiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
