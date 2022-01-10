package knet

import "github.com/KarlvenK/kinx/kiface"

type Request struct {
	//build connect
	conn kiface.IConnection
	//data of request from client
	data []byte
}

func (r *Request) GetConnection() kiface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
