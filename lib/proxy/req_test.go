package proxy

import (
	"fmt"
	"github.com/leicc520/go-orm/log"
	"testing"
)

func TestName(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	link := "https://www.amazon.com/sp?seller=A2VTI8ZSUTN12G&language=en_US"
	//link  = "https://www.amazon.com/sp?seller=A2Y12R7CMI5RDY&language=en_US"
	//link  = "http://test.loc/index.php?aaa=xxx"
	client := NewHttpRequest()
	//client.AddHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	client.AddHeader("accept-encoding", "gzip, deflate, br")
	client.AddHeader("accept-language", "zh-CN,zh;q=0.9")
	client.AddHeader("upgrade-insecure-requests", "1")
	/*
	client.AddHeader("origin","https://www.amazon.com")
	client.AddHeader("authority", "www.amazon.com")
	client.AddHeader("referer", "https://www.amazon.com/")
	client.AddHeader("user-agent", RandUserAgent())
*/
	//获取店铺信息的时候 需要使用session-id
	for i := 0; i < 2; i++ {
		result, err := client.Request(link, nil, "GET")
		sessid :=  client.GetJarCookie("https://www.amazon.com/", "session-id")
		fmt.Println(i, err, sessid, result, "=======")
		if err == nil {
			break
		}
		client.AddHeader("cookie", "session-id="+sessid)
	}
}
