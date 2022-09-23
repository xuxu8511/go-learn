package xnet

import "testing"

func TestNewClient(t *testing.T) {
	for i := 0; i <= 100; i++ {
		go NewClient("124.223.83.101", 9999)
		//go NewClient("127.0.0.1", 9999)
	}
	select {}
}
