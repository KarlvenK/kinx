package knet

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/utils"
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
	//current server's msgHandler
	MsgHandler kiface.IMsgHandle
	//connManager of curr server
	ConnMgr kiface.IConnManager
	//hook func when start
	OnConnStart func(conn kiface.IConnection)
	//hook func when stop
	OnConnStop func(conn kiface.IConnection)
}

/*

func callback(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

*/

func (s *Server) Start() {
	fmt.Printf("[kinx]Server name: %s, listenner at IP: %s, Port: %d is starting...\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	go func() {
		//start workerpool
		s.MsgHandler.StartWorkerPool()

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

		var cid uint32
		cid = 0

		//block client connection handle client's work
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				fmt.Println("Too many Connections MaxConn = ", utils.GlobalObject.MaxConn)
				_ = conn.Close()
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[stop]Kinx server name ", s.Name)
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	//start server services
	s.Start()

	//todo

	//block
	select {}

}

func (s *Server) AddRouter(msgID uint32, router kiface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router succ!")
}

// NewServer init Server module
func NewServer() kiface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

func (s *Server) GetConnMgr() kiface.IConnManager {
	return s.ConnMgr
}

func (s *Server) SetOnConnStart(hookFunc func(conn kiface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(conn kiface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn kiface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("--->Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn kiface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("--->Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}
