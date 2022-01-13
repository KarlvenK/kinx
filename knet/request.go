package knet

import "github.com/KarlvenK/kinx/kiface"

type Request struct {
	//build connect
	conn kiface.IConnection
	//data of request from client
	msg kiface.IMessage
}

func (r *Request) GetConnection() kiface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
