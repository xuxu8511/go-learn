package xnet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Router interface {
	PreHandle(request Request)
	Handle(request Request)
	PostHandle(request Request)
}

type BaseRouter struct{}

func (router BaseRouter) PreHandle(request Request) {
	//fmt.Println("Router PreHandle")
	request.GetConnection().SetProperty("abc", "111")
}

func (router BaseRouter) Handle(request Request) {
	//fmt.Println("Router Handle")
	//_ := request.GetConnection()
	data := request.GetData()
	v, _ := request.GetConnection().GetProperty("abc")
	log.Info("property test, v:", v)
	log.Info(fmt.Sprintf("receive client msg:%s", data))

	//if _, err := connection.Send2(data); err != nil {
	//	fmt.Println("send failed, err", err)
	//}
}

func (router BaseRouter) PostHandle(request Request) {
	//fmt.Println("Router PostHandle")
}
