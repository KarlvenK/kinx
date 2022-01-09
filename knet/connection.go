package knet

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"net"
)

type Connection struct {
	//socket TCP
	Conn *net.TCPConn

	//Conn ID
	ConnID uint32

	//current stat
	isClosed bool

	//API that current conn binding
	handleAPI kiface.HandleFunc

	//notify current conn to exit
	ExitChan chan bool
}

// NewConnection Init connection module
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi kiface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callbackApi,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.Conn)
	go c.StartReader()
	//todo! run current conn write work

}

func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//Read the data from client
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}

		//run HandleAPI
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID = ", c.ConnID, "handle is err", err)
			break
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//close socket
	c.Conn.Close()

	close(c.ExitChan)

}

func (c *Connection) GetConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
