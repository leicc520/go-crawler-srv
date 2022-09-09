package proxy

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
	"math/rand"
	"net/http"
	"testing"
)

func TestStatistic(t *testing.T) {
	str := "redis://:@127.0.0.1:6379/1"
	lib.InitRedis(str)
	proxy := []ProxySt{{Proxy: "127.0.0.1:11", Status: 1},{Proxy: "127.0.0.1:22", Status: 1},{Proxy: "127.0.0.1:33", Status: 1}}

	ss := NewStatistic(proxy, lib.Redis)

	for {
		idx := rand.Int() % len(proxy)
		host:= fmt.Sprintf("demo.%d", rand.Int()%100)
		req, _ := http.NewRequest("GET", "http://"+host+"/demo", nil)
		sp := http.Response{StatusCode: http.StatusOK}
		ss.Report(idx, req, &sp)
	}

	ch := make(chan int)
	<-ch
}