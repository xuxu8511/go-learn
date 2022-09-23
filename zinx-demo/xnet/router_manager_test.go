package xnet

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	testChan := make(chan int, 1000000)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case result := <-testChan:
					fmt.Println(result)
				}
			}
		}()
	}

	go func() {
		for {
			testChan <- rand.Int()
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe(":8787", nil)

	select {}
}
