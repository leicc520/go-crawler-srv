package channal

import (
	"github.com/leicc520/go-orm/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type EasyGoSt struct {
}

//获取代理IP 数据资料信息
func (s EasyGoSt) GetProxy(proto string) string {
	query := url.Values{}
	query.Set("count", "1")
	query.Set("neek", "48389")
	query.Set("yys", "0")
	query.Set("type", "1")
	query.Set("mr", "1")
	query.Set("sb", "")
	query.Set("sep", "1")
	switch proto {
	case PROXY_HTTPS:
		query.Set("port", "2")
	case PROXY_HTTP:
		query.Set("port", "1")
	default:
		query.Set("port", "11")
	}
	sp, err := (&http.Client{}).Get(EASY_GO_PROXY+"/getip?"+query.Encode())
	if err != nil || sp == nil || sp.StatusCode != http.StatusOK {
		log.Write(-1, "easy proxy error ", err)
		return ""
	}
	defer sp.Body.Close()
	body, _ := ioutil.ReadAll(sp.Body)
	bodyStr := strings.TrimSpace(string(body))
	if ok, err := regexp.MatchString("\\:[\\d]+$", bodyStr); !ok || err != nil {
		log.Write(-1, "easy proxy error ", bodyStr)
		return ""
	}
	return bodyStr
}
