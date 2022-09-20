package xnet

import log "github.com/sirupsen/logrus"

type RouterManager struct {
	routers map[uint32]Router
}

func NewRouterManager() *RouterManager {
	return &RouterManager{
		routers: make(map[uint32]Router),
	}
}

func (rm *RouterManager) DoRouter(request Request) {
	rt, ok := rm.routers[request.GetData().GetMsgId()]
	if !ok {
		log.Error("not find router: %#", request.GetData().GetMsgId())
		return
	}

	rt.PreHandle(request)
	rt.Handle(request)
	rt.PostHandle(request)
}

func (rm *RouterManager) AddRouter(id uint32, router Router) {
	rm.routers[id] = router
}
