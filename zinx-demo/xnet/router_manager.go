package xnet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type RouterManager struct {
	routers       map[uint32]Router
	taskWaitQueue []chan Request
	workerSize    uint32
	maxWaitSize   uint32
}

func NewRouterManager(maxWaitSize, workerSize uint32) *RouterManager {
	return &RouterManager{
		routers:       make(map[uint32]Router),
		workerSize:    workerSize,
		maxWaitSize:   maxWaitSize,
		taskWaitQueue: make([]chan Request, workerSize),
	}
}

func (rm *RouterManager) DoRouter(request Request) {
	log.Info("DoRouter.1...")
	intN := rand.Intn(int(rm.workerSize))
	rm.taskWaitQueue[intN] <- request
	log.Info("DoRouter.2...")
}

func (rm *RouterManager) AddRouter(id uint32, router Router) {
	rm.routers[id] = router
}

func (rm *RouterManager) StartOneWorker(taskQueue chan Request) {
	log.Info("one work start")
	defer func() {
		log.Info("worker exit")
	}()

	for {
		select {
		case request := <-taskQueue:
			log.Info(fmt.Sprintf("taskQueue request:%#v", request))
			rt, ok := rm.routers[request.GetData().GetMsgId()]
			if !ok {
				log.Error(fmt.Sprintf("not find router: %#v", request.GetData().GetMsgId()))
				return
			}

			rt.PreHandle(request)
			rt.Handle(request)
			rt.PostHandle(request)
		}
	}
}

func (rm *RouterManager) StartWorker() {
	for i := 0; i < int(rm.workerSize); i++ {
		rm.taskWaitQueue[i] = make(chan Request, rm.maxWaitSize)
		go rm.StartOneWorker(rm.taskWaitQueue[i])
	}

	go func() {
		newTimer := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-newTimer.C:
				for i := 0; i < int(rm.workerSize); i++ {
					log.Info(fmt.Sprintf("taskWaitSize i:%d remain:%d", i, len(rm.taskWaitQueue[i])))
				}
			}
		}
	}()
}
