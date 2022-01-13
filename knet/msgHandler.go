package knet

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"strconv"
)

type MsgHandle struct {
	//store every handler of MsgID
	APis map[uint32]kiface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		APis: make(map[uint32]kiface.IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request kiface.IRequest) {
	handle, ok := mh.APis[request.GetMsgID()]
	if !ok {
		fmt.Println("api MsgID = ", request.GetMsgID(), " is not found. Need to register")
	} else {
		handle.PreHandle(request)
		handle.Handle(request)
		handle.PostHandle(request)
	}
}

func (mh *MsgHandle) AddRouter(msgID uint32, router kiface.IRouter) {
	if _, ok := mh.APis[msgID]; ok {
		panic("repeat api, msgID = " + strconv.Itoa(int(msgID)))
	} else {
		mh.APis[msgID] = router
		fmt.Println("add api MsgID = ", msgID, "succ")
	}
}
