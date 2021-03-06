package knet

import (
	"errors"
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/utils"
	"io"
	"net"
	"sync"
)

type Connection struct {
	//server that curr conn belongs to
	TcpServer kiface.IServer

	//socket TCP
	Conn *net.TCPConn

	//Conn ID
	ConnID uint32

	//current stat
	isClosed bool

	//notify current conn to exit
	ExitChan chan bool

	msgChan chan []byte

	MsgHandler kiface.IMsgHandle

	property map[string]interface{}

	propertyLock sync.RWMutex
}

// NewConnection Init connection module
func NewConnection(server kiface.IServer, conn *net.TCPConn, connID uint32, msgHandler kiface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
		property:   make(map[string]interface{}),
	}
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.Conn)
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallOnConnStart(c)
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer goroutine is running...]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader goroutine is running...]")
	defer fmt.Println("connID = ", c.ConnID, "[Reader is exit!], remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
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

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//find the corresponding router of conn
			//do the handle func
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnStop(c)

	//close socket
	_ = c.Conn.Close()
	c.ExitChan <- true

	c.TcpServer.GetConnMgr().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)

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

	c.msgChan <- binaryMsg

	return nil
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
