package proxy

import (
	"fmt"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	aa := time.Second*3
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second*1)
			ch <-i
		}
	}()
	time.AfterFunc(time.Second*3, func() {
		fmt.Println("=====after")
	})
	timer := time.After(aa)
	for {
		select {
			case <-timer:
				fmt.Println("====OK")
				timer = time.After(aa)
			case bb:=<-ch:
				fmt.Println("==========", bb)
		}
	}
}
