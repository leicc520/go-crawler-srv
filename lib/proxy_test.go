package lib

import (
	"fmt"
	proxy2 "github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-crawler-srv/lib/proxy/channal"
	"math/rand"
	"net/http"
	"testing"
	"time"
)


func TestProxy(t *testing.T) {
	ss := channal.EasyGoSt{}
	str := ss.GetProxy(channal.PROXY_SOCK5)
	fmt.Println(str)
}

func TestStatistic(t *testing.T) {
	str := "redis://:@127.0.0.1:6379/1"
	InitRedis(str)
	proxy := []proxy2.ProxySt{{Proxy: "127.0.0.1:11", Status: 1},{Proxy: "127.0.0.1:22", Status: 1},{Proxy: "127.0.0.1:33", Status: 1}}
	ss := proxy2.NewStatistic(proxy, Redis)
	go func() {
		status := []int{200,302,304,404,403,500,502,504}
		for {
			idx := rand.Int() % len(proxy)
			host:= fmt.Sprintf("demo.%d", rand.Int()%100)
			req, _ := http.NewRequest("GET", "http://"+host+"/demo", nil)
			mystatus := status[rand.Int()%len(status)]
			sp := http.Response{StatusCode: mystatus}
			ss.Report(idx, req.Host, sp.StatusCode)
			time.Sleep(time.Second)
		}
	}()

	for  {
		i, str := ss.Proxy()
		time.Sleep(time.Second)
		fmt.Println("str:", i, str)
	}

}