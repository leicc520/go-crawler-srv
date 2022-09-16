package spider

import (
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-orm/cache"
)

var (
	monitor *proxy.Monitor = nil
	mCache  cache.Cacher   = nil
)

type SpiderSt struct {


}

func (s *SpiderSt) Init() {

}
