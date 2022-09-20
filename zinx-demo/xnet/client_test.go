package xnet

import "testing"

func TestNewClient(t *testing.T) {
	NewClient("0.0.0.0", 9999)
}
