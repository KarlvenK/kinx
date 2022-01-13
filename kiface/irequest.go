package kiface

type IRequest interface {
	// GetConnection get current connecnt
	GetConnection() IConnection
	// GetData get requestion data
	GetData() []byte

	GetMsgID() uint32
}
