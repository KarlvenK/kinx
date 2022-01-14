package kiface

type IServer interface {
	// Start start server
	Start()
	// Stop stop server
	Stop()
	// Serve run server
	Serve()
	// AddRouter add router func for current service so that client can use
	AddRouter(msgID uint32, router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(conn IConnection))

	SetOnConnStop(func(conn IConnection))

	CallOnConnStart(conn IConnection)

	CallOnConnStop(conn IConnection)
}
