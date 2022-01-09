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
	// Send data to client
	Send(data []byte) error
}

// HandleFunc define a func handle work
type HandleFunc func(*net.TCPConn, []byte, int) error
