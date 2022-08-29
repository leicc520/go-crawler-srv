package proxy

import (
	"fmt"
	"github.com/leicc520/go-orm/log"
	"testing"
)

func TestName(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	link := "http://10.71.32.65:22336/s?me=A29LJNDOV7ZB72&language=en_US"
	client := NewHttpRequest()
	client.AddHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	client.AddHeader("accept-encoding", "gzip, deflate, br")
	client.AddHeader("accept-language", "en")
	client.AddHeader("upgrade-insecure-requests", "1")
	client.AddHeader("origin","https://www.amazon.com")
	client.AddHeader("referer", "https://www.amazon.com/")
	client.AddHeader("user-agent", RandUserAgent())
	result, err := client.Request(link, nil, "GET")
	fmt.Println(result, err, "=======")
}
