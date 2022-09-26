package xnet

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"sync"
)

type Connection interface {
	Start()

	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	Send(data []byte) (int, error)

	Send2(message Message) (int, error)

	SetProperty(string, interface{})

	GetProperty(string) (interface{}, bool)
}

type ConnectionImpl struct {
	Conn         *net.TCPConn
	ConnID       uint32
	IsClosed     bool
	ExitChan     chan bool        //通知连接退出的channel
	MsgChan      chan interface{} //用于读写goroutine的通信，同步机制
	RM           *RouterManager
	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func (c *ConnectionImpl) SetProperty(s string, i interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[s] = i
}

func (c *ConnectionImpl) GetProperty(s string) (interface{}, bool) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	v, ok := c.property[s]
	if !ok {
		return nil, false
	}
	return v, true
}

func NewConnectionImpl(conn *net.TCPConn, connID uint32, rm *RouterManager) Connection {
	c := &ConnectionImpl{
		Conn:     conn,
		ConnID:   connID,
		IsClosed: false,
		RM:       rm,
		ExitChan: make(chan bool, 1),
		MsgChan:  make(chan interface{}),
		property: make(map[string]interface{}),
	}

	return c
}

func (c *ConnectionImpl) StartReader() {
	log.Info("read goroutine is running")
	defer log.Info("ConnID=", c.ConnID, " Reader goroutine exit, addr:", c.RemoteAddr().String())
	defer c.Stop()

	for {
		codec := NewCodeCImpl()
		headData := make([]byte, codec.GetHeaderLen())

		//c.Conn.Read() 会导致内存溢出，使用io.ReadFull则不会。
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			log.Error("recv buf error", err)
			return
		}

		message, err := codec.Decoder(headData)
		if err != nil {
			log.Error("decoder error:", err)
			return
		}
		log.Info("=====================>", message.GetMsgLength(), message.GetMsgId())

		data := make([]byte, message.GetMsgLength())
		if _, err = io.ReadFull(c.Conn, data); err != nil {
			log.Error("read msg error", err)
			return
		}
		message.SetData(data)

		r := &RequestImpl{
			conn: c,
			msg:  message,
		}

		c.RM.DoRouter(r)
	}
}

/*
写消息goroutine，专门发送消息给客户端
*/
func (c *ConnectionImpl) StartWriter() {
	log.Info("writer goroutine is running")
	defer log.Info("ConnID=", c.ConnID, " writer goroutine exit, addr:", c.RemoteAddr().String())
	for {
		select {
		case msg := <-c.MsgChan:
			log.Info("writer goroutine recv msg:{}", msg)
			if _, err := c.Conn.Write(msg.([]byte)); err != nil {
				log.Error("err:{}", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *ConnectionImpl) Start() {
	log.Info("Conn StartGoroutine Start, ConnIdD=", c.ConnID)

	go c.StartReader()
	go c.StartWriter()

}

func (c *ConnectionImpl) Stop() {
	if c.IsClosed == true {
		return
	}

	c.IsClosed = true

	c.Conn.Close()

	c.ExitChan <- true
	close(c.ExitChan)
	close(c.MsgChan)
}

func (c *ConnectionImpl) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *ConnectionImpl) GetConnID() uint32 {
	return c.ConnID
}

func (c *ConnectionImpl) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *ConnectionImpl) Send(data []byte) (int, error) {
	c.MsgChan <- data
	return 0, nil
}

func (c *ConnectionImpl) Send2(message Message) (int, error) {
	codec := NewCodeCImpl()
	encoder, err := codec.Encoder(message)
	if err != nil {
		log.Error("encoder err", err)
		return 0, err
	}
	return c.Send(encoder)
}
