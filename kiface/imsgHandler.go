package kiface

type IMsgHandle interface {
	// DoMsgHandler run bingding router func
	DoMsgHandler(request IRequest)
	// AddRouter add handle func
	AddRouter(msgID uint32, router IRouter)

	StartWorkerPool()

	SendMsgToTaskQueue(request IRequest)
}
