package knet

import (
	"errors"
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"io"
	"net"
)

type Connection struct {
	//socket TCP
	Conn *net.TCPConn

	//Conn ID
	ConnID uint32

	//current stat
	isClosed bool

	//notify current conn to exit
	ExitChan chan bool

	MsgHandler kiface.IMsgHandle
}

// NewConnection Init connection module
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler kiface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
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
		/*
			buf := make([]byte, utils.GlobalObject.MaxPackageSize)
			_, err := c.Conn.Read(buf)
			if err != nil {
				fmt.Println("recv buf err ", err)
				continue
			}*/

		//do what we do in datapack_test.go
		dp := NewDataPack()
		msgHead := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), msgHead)
		if err != nil {
			fmt.Println("read msgHead err", err)
			break
		}

		msg, err := dp.Unpack(msgHead)
		if err != nil {
			fmt.Println("unpack err", err)
			break
		}
		if msg.GetMsgLen() > 0 {
			data := make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read msg data err", err)
				break
			}
			msg.SetData(data)
		}

		req := Request{
			conn: c,
			msg:  msg,
		}
		//find the corresponding router of conn
		//do the handle func
		go c.MsgHandler.DoMsgHandler(&req)
	}
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//close socket
	_ = c.Conn.Close()

	close(c.ExitChan)

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}

	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id = ", msgId)
		return errors.New("pack error msg")
	}

	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write msg id", msgId, " error :", err)
		return errors.New("conn Write error")
	}

	return nil
}
