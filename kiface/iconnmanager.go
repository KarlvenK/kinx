package kiface

type IConnManager interface {
	Add(conn IConnection)                   //add conn
	Remove(conn IConnection)                //remove conn
	Get(connID uint32) (IConnection, error) //get conn according to id
	Len() int                               //get current cnt of conns
	ClearConn()                             //clear and stop all conn
}
