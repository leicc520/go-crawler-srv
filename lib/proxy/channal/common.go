package channal

const (
	EASY_GO_PROXY = "http://zltiqu.pyhttp.taolop.com"

	PROXY_SOCK5   = "tcp"
	PROXY_HTTPS   = "https"
	PROXY_HTTP    = "http"
)

type IFProxy interface {
	GetProxy(proto string) string
}

/******************************************************
  把系统支持的代理放写到这里，进行管理 需要加白明名单才能请求
 */
