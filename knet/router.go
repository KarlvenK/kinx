package knet

import "github.com/KarlvenK/kinx/kiface"

//实现Router时，先嵌入这个BaseRouter基类，然后根据需求重写方法

type BaseRouter struct{}

//之所以BaseRouter的方法都为空，因为有些router不希望又PreHandle PostHandle方法
//所以Router全部继承BaseRouter的好处就是，不需要实现PreHandle PostHandle， 也可以实例化

func (b *BaseRouter) PreHandle(request kiface.IRequest)  {}
func (b *BaseRouter) Handle(request kiface.IRequest)     {}
func (b *BaseRouter) PostHandle(request kiface.IRequest) {}
