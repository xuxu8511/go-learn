package xnet

import (
	"fmt"
	"net"
	"time"
)

func NewClient(ip string, port int32) {

	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Errorf("dial error: %+v\n", err)
		return
	}

	for {
		time.Sleep(1 * time.Second)
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
	}

}
