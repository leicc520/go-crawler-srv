package spider

import (
	"errors"
	"net/url"

	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
)

const (
	SpiderDataExpire = 86400
)

/********************************************************************
	请求业务的封装，获取到数据之后写缓存，然后返回
 */
type BaseRequest struct {
	Url 	string 				`json:"url"`
	Cookie 	string 				`json:"cookie"`
	Agent   string 				`json:"agent"`
	Method  string				`json:"method"`
	Params  string              `json:"params"`
	Header  map[string]string 	`json:"header"`
}

//获取缓存策略的key
func (r BaseRequest) CacheKey(uri *url.URL) string {
	return uri.Host+"@"+lib.Md5Str(uri.String())
}

//通过缓存获取数据
func (r BaseRequest) CacheGet(uri *url.URL) (ckey, result string) {
	ckey    = r.CacheKey(uri)
	cResult := mCache.Get(ckey)
	if cResult != nil {
		if lStr, ok := cResult.(string); ok {
			result   = lStr
		}
	}
	return
}

//发起网络请求，爬取业务数据资料信息
func (r BaseRequest) Do() (string, error) {
	uri, err := url.Parse(r.Url)
	if err != nil {
		return "", errors.New("地址格式错误:"+r.Url)
	}
	ckey, result := r.CacheGet(uri)
	if len(result) > 0 {//缓存已经存在
		return result, err
	}
	client := proxy.NewHttpRequest().SetMonitor(monitor)
	if r.Header == nil {//初始化
		r.Header = make(map[string]string)
	}
	if len(r.Agent) > 0 {
		r.Header["user-agent"] = r.Agent
	}
	if len(r.Cookie) > 0 {
		baseUrl := uri.Scheme + "://"+uri.Host + "/"
		client.SetCookie(baseUrl, r.Cookie)
	}
	if len(r.Header) > 0 {//设置请求头信息
		client.SetHeader(r.Header)
	}
	result, err = client.Request(r.Url, []byte(r.Params), r.Method)
	if err == nil && len(result) > 0 {//爬取到内容了
		mCache.Set(ckey, result, SpiderDataExpire)
	}
	return result, err
}
