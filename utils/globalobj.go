package utils

import (
	"encoding/json"
	"github.com/KarlvenK/kinx/kiface"
	"io/ioutil"
)

//store all parameters of kinx
//definee by kinx.json

type GlobalObj struct {
	TcpServer kiface.IServer //current kinx global server obj
	Host      string         //current server host ip
	TcpPort   int            //port
	Name      string         //server name

	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

//Reload load parameters from knix.json
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("config/kinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//init initialize current GlobalObj automatically
func init() {
	//default
	GlobalObject = &GlobalObj{
		Name:           "kinxServerApp",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//load from knix.json
	GlobalObject.Reload()
}
