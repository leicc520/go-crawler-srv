package parser

import (
	"errors"
	"github.com/leicc520/go-orm/cache"
	"net/url"
	"regexp"
	"strings"

	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
)

const (
	SpiderDataExpire = 86400
)

var (
	monitor *proxy.Monitor = nil
	mCache  cache.Cacher   = nil
	errUnknownPage = errors.New("unknown page")
)

/********************************************************************
请求业务的封装，获取到数据之后写缓存，然后返回
*/
type BaseRequest struct {
	Url 		string 				`json:"url"     yaml:"url"`
	RegUrl  	[]string            `json:"reg_url" yaml:"reg_url"`
	RegMatch 	[]string 			`json:"reg_match" yaml:"reg_match"`
	Method  	string				`json:"method"  yaml:"method"`
	Params  	string              `json:"params"  yaml:"params"`
	Header  	map[string]string 	`json:"headers" yaml:"headers"`
}

//注入缓存以及代理监控
func Inject(sCache cache.Cacher, sMonitor *proxy.Monitor) {
	if sCache != nil {
		mCache = sCache
	}
	if sMonitor != nil {
		monitor= sMonitor
	}
}

//获取缓存策略的key
func (r *BaseRequest) CacheKey(uri *url.URL) string {
	return uri.Host+"@"+lib.Md5Str(uri.String())
}

//通过缓存获取数据
func (r *BaseRequest) CacheGet(uri *url.URL) (ckey, result string) {
	ckey    = r.CacheKey(uri)
	cResult := mCache.Get(ckey)
	if cResult != nil {
		if lStr, ok := cResult.(string); ok {
			result   = lStr
		}
	}
	return
}

//验证请求的地址是否和当前任务匹配
func (r *BaseRequest) isRegUrl() bool {
	if len(r.RegUrl) > 0 {
		for _, regStr := range r.RegUrl {
			ok, err := regexp.MatchString(regStr, r.Url)
			if ok && err == nil {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

//检测获取的内容是否和预期的一致
func (r *BaseRequest) isRegMatch(result string) bool {
	if len(r.RegMatch) > 0 {
		for _, regStr := range r.RegMatch {
			if ok, err := regexp.MatchString(regStr, result); ok && err == nil {
				return true
			}
		}
	} else {
		return true
	}
	return false
}

//发起网络请求，爬取业务数据资料信息
func (r *BaseRequest) Do(link string) (string, error) {
	if len(link) > 0 && strings.HasPrefix(link, "http") {
		r.Url = link
	}
	uri, err := url.Parse(r.Url)
	if err != nil {
		return "", errors.New("地址格式错误:"+r.Url)
	}
	ckey, result := r.CacheGet(uri)
	if len(result) > 0 {//缓存已经存在
		return result, err
	}
	if !r.isRegUrl() {//地址不匹配的情况
		lib.LogActionOnce("RegUrl", 300, r.RegUrl,  r.Url)
		return "", errors.New("地址模式不匹配:"+r.Url)
	}
	client := proxy.NewHttpRequest().SetMonitor(monitor)
	if r.Header != nil && len(r.Header) > 0 {//设置请求头信息
		client.SetHeader(r.Header)
	}
	result, err = client.Request(r.Url, []byte(r.Params), r.Method)
	if err == nil && len(result) > 0 && !r.isRegMatch(result) {
		err = errUnknownPage
	}
	//过滤一下空格处理逻辑
	result = lib.NormalizeSpace(result)
	if err == nil && len(result) > 0 {//爬取到内容了
		mCache.Set(ckey, result, SpiderDataExpire)
	}
	return result, err
}

