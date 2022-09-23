package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	. "zinx-demo/global"
	. "zinx-demo/xnet"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Info(err)
			fmt.Printf("%+v", err)
		}
	}()

	s := NewServer("tcpserver", "tcp4", G_ZinxConfig.Server.Ip, int32(G_ZinxConfig.Server.Port))
	routerManager := NewRouterManager(G_ZinxConfig.Server.MaxWaitSize, G_ZinxConfig.Server.HandleWorkerSize)
	routerManager.AddRouter(1, &BaseRouter{})
	routerManager.AddRouter(2, &BaseRouter{})
	s.AddRouterManager(routerManager)
	go s.Serve()

	go http.ListenAndServe(":8080", nil)
	select {}
}
