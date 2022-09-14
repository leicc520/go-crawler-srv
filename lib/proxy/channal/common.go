package channal

import "sync"

const (
	EASY_GO_PROXY = "http://zltiqu.pyhttp.taolop.com"
	SKY_GO_PROXY  = "http://api.tianqiip.com"
	PROXY_SOCK5   = "tcp"
	PROXY_HTTPS   = "https"
	PROXY_HTTP    = "http"
	PROXY_CHANNEL_EASYGO = "easygo"
	PROXY_CHANNEL_SKYGO  = "skygo"
)

type IFProxy interface {
	GetProxy(proto string) string
}

var (
	proxyOnce  = sync.Once{}
	proxyDriver map[string]IFProxy = nil
)

//代理注册到注册器当中
func proxyRegister(name string, ifProxy IFProxy)  {
	proxyOnce.Do(func() {//初始化逻辑
		proxyDriver = make(map[string]IFProxy)
	})
	proxyDriver[name] = ifProxy
}

//获取代理IP数据资料信息
func GetProxy(name,proto string) string  {
	if proxyDriver == nil {
		return ""
	}
	if s, ok := proxyDriver[name]; ok {
		return s.GetProxy(proto)
	}
	return ""
}
/******************************************************
  把系统支持的代理放写到这里，进行管理 需要加白明名单才能请求
 */
