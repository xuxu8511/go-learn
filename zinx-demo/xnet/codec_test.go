package xnet

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"testing"
)

func TestCodeCImpl_Decoder(t *testing.T) {

	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	go func() {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("accept error:", err2)
			return
		}

		go func(conn net.Conn) {
			dp := NewCodeCImpl()
			for {
				headData := make([]byte, dp.GetHeaderLen())
				_, err3 := io.ReadFull(conn, headData)
				if err3 != nil {
					fmt.Println("read head error", err3)
					break
				}

				message, err3 := dp.Decoder(headData)
				if err3 != nil {
					fmt.Println("decoder head error", err3)
					break
				}
				log.Info("%#", message)
				message.GetMsgLength()
				data := make([]byte, message.GetMsgLength())
				_, err3 = io.ReadFull(conn, data)
				if err3 != nil {
					fmt.Println("read data error", err3)
					return
				}
				message.SetData(data)
				log.Info(fmt.Sprintf("recv message: %v", message))
			}
		}(conn)
	}()

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("dial conn error:", err)
		return
	}

	dp := NewCodeCImpl()

	msg1 := &MessageImpl{
		Id:     1,
		Length: 4,
		Data:   []byte{'x', 'n', 'e', 't'},
	}
	bytes1, err := dp.Encoder(msg1)
	if err != nil {
		fmt.Println("encoder err:", err)
		return
	}

	msg2 := &MessageImpl{
		Id:     2,
		Length: 4,
		Data:   []byte{'1', '2', '3', '4'},
	}
	bytes2, err2 := dp.Encoder(msg2)
	if err2 != nil {
		fmt.Println("encoder err:", err)
		return
	}
	bytes1 = append(bytes1, bytes2...)

	conn.Write(bytes1)

	select {}

}
