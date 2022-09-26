package xnet

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"testing"
)

func TestTimer(t *testing.T) {
	for i := 0; i <= 100; i++ {
		intN := rand.Intn(int(20))
		fmt.Println(intN)
	}

}
