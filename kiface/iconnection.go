package kiface

import "net"

type IConnection interface {
	// Start connection
	Start()
	// Stop connection
	Stop()
	// GetTCPConnection get socket conn that current connection binding
	GetTCPConnection() *net.TCPConn
	// GetConnID get cur conn ID
	GetConnID() uint32
	// RemoteAddr get remote client TCP stat IP port
	RemoteAddr() net.Addr
	// SendMsg data to client
	SendMsg(uint32, []byte) error

	SetProperty(key string, value interface{})

	GetProperty(key string) (interface{}, error)

	RemoveProperty(key string)
}
