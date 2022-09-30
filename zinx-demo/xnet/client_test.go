package xnet

import "testing"

func TestNewClient(t *testing.T) {
	for i := 0; i <= 5000; i++ {
		go NewClient("192.168.100.148", 6666)
		//go NewClient("127.0.0.1", 9999)
	}
	select {}
}
