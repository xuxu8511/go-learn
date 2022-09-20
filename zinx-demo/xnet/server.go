package xnet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type Server interface {
	Start()
	Stop()
	Serve()
	AddRouterManager(manager *RouterManager)
}

type ServerImpl struct {
	Name      string
	IPVersion string
	Ip        string
	Port      int32
	RM        *RouterManager
}

func NewServer(name, ipVersion, ip string, port int32) Server {
	s := &ServerImpl{
		Name:      name,
		IPVersion: ipVersion,
		Ip:        ip,
		Port:      port,
		RM:        nil,
	}

	return s
}

func (s *ServerImpl) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
		if err != nil {
			fmt.Errorf("resolve ip addr:%+v\n", err)
			return
		}

		tcpListener, err := net.ListenTCP("tcp4", addr)
		if err != nil {
			fmt.Errorf("ListenTCP error:%+v\n", err)
			return
		}

		log.Info(fmt.Sprintf("tcp server[%s] init complete", fmt.Sprintf("%s:%d", s.Ip, s.Port)))

		var cid uint32 = 0
		for {
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error", err)
				continue
			}

			connection := NewConnectionImpl(tcpConn, cid, s.RM)
			cid++

			go connection.Start()
		}
	}()
}

func (s *ServerImpl) Stop() {
	fmt.Printf("server stoped")
}

func (s *ServerImpl) Serve() {
	s.Start()

	log.Info("===============server init complete======================")

	//阻塞
	//time.Sleep(1 * time.Hour)
	select {}
}

func (s *ServerImpl) AddRouterManager(manager *RouterManager) {
	s.RM = manager
}
