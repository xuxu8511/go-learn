package main

import (
	"zinx-demo/global"
	. "zinx-demo/xnet"
)

func main() {
	s := NewServer("tcpserver", "tcp4", global.G_ZinxConfig.Server.Ip, int32(global.G_ZinxConfig.Server.Port))

	routerManager := NewRouterManager()
	routerManager.AddRouter(1, &BaseRouter{})
	routerManager.AddRouter(2, &BaseRouter{})
	s.AddRouterManager(routerManager)
	s.Serve()
}
