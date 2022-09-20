package xnet

import "fmt"

type PingRouter struct {
}

func (p *PingRouter) PreHandle(request Request) {
}

func (p *PingRouter) Handle(request Request) {
	connection := request.GetConnection()
	data := request.GetData()
	fmt.Println("receive msg:", data)

	if _, err := connection.Send([]byte("pong")); err != nil {
		fmt.Println("send failed, err:", err)
	}
}

func (p *PingRouter) PostHandle(request Request) {
}
