package knet

import "github.com/KarlvenK/kinx/kiface"

//Server impl IServer interface
//defines a Server mmodule
type Server struct {
	//server name
	Name string
	//binding ip version
	IPVersion string
	//listened IP
	IP string
	//listened port
	Port string
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}

// NewServer init Server module
func NewServer(name string) kiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      "8999",
	}
	return s
}
