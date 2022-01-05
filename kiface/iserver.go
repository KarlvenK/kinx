package kiface

type IServer interface {
	//start server
	Start()
	//stop server
	Stop()
	//run server
	Serve()
}
