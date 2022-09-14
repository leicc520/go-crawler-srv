package channal

import (
	"github.com/leicc520/go-orm/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type SkyStartGoSt struct {
}

//获取代理IP 数据资料信息
func (s SkyStartGoSt) GetProxy(proto string) string {
	query := url.Values{}
	query.Set("secret", "1lybseonkysemhod")
	query.Set("num", "1")
	query.Set("type", "txt")
	query.Set("time", "15")
	query.Set("sign", "ca2c72d5b18a53b440f54e80d59f773f")
	switch proto {
	case PROXY_HTTPS:
		query.Set("port", "2")
	case PROXY_HTTP:
		query.Set("port", "1")
	default:
		query.Set("port", "3")
	}
	sp, err := (&http.Client{}).Get(SKY_GO_PROXY+"/getip?"+query.Encode())
	if err != nil || sp == nil || sp.StatusCode != http.StatusOK {
		log.Write(-1, "sky proxy error ", err)
		return ""
	}
	defer sp.Body.Close()
	body, _ := ioutil.ReadAll(sp.Body)
	bodyStr := strings.TrimSpace(string(body))
	if ok, err := regexp.MatchString("\\:[\\d]+$", bodyStr); !ok || err != nil {
		log.Write(-1, "sky proxy error ", bodyStr)
		return ""
	}
	return bodyStr
}
