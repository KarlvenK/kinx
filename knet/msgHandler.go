package knet

import (
	"fmt"
	"github.com/KarlvenK/kinx/kiface"
	"github.com/KarlvenK/kinx/utils"
	"strconv"
)

type MsgHandle struct {
	//store every handler of MsgID
	APis map[uint32]kiface.IRouter

	//message queue
	TaskQueue []chan kiface.IRequest

	//cnt of workers of workerpool
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		APis:           make(map[uint32]kiface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan kiface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

// StartWorkerPool start a workerpool
func (mh *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan kiface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// StartOneWorker start a worker process
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan kiface.IRequest) {
	fmt.Println("WorkerID = ", workerID, " is running...")
	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandle) SendMsgToTaskQueue(request kiface.IRequest) {
	//distribute to worker
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		" request MsgID = ", request.GetMsgID(),
		" to workerID = ", workerID)

	mh.TaskQueue[workerID] <- request
}
