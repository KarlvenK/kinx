package kiface

type IServer interface {
	// Start start server
	Start()
	// Stop stop server
	Stop()
	// Serve run server
	Serve()
}
