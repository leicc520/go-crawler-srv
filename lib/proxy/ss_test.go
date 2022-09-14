package proxy

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib/proxy/channal"
	"testing"
	"time"
)

func TestProxyChannel(t *testing.T) {
	ss := ProxySt{Url: "", Proxy: channal.PROXY_CHANNEL_EASYGO, Status: 1}
	ss.CutProxy()
	fmt.Println(ss)
}

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
