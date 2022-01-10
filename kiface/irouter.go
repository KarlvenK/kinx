package kiface

type IRouter interface {
	// PreHandle before conn hook
	PreHandle(request IRequest)
	// Handle handle conn hook
	Handle(request IRequest)
	// PostHandle after conn hook
	PostHandle(request IRequest)
}
