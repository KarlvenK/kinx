package kiface

/*
pack depack module
orienated to the stream of TCP conn
*/

type IDataPack interface {
	GetHeadLen() uint32

	Pack(IMessage) ([]byte, error)

	Unpack([]byte) (IMessage, error)
}
