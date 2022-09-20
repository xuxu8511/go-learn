package xnet

import (
	"fmt"
	"testing"
)

type Abc struct {
	v int32
}

func NewAbc(v int32) *Abc {
	return &Abc{v: v}
}

func (a Abc) change() {
	a.v = 1
}

func (a *Abc) change2() {
	a.v = 1
}

func TestNewServer(t *testing.T) {
}

func TestPrint(t *testing.T) {
	fmt.Printf("1111111111111111")
}
