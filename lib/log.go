package lib

import (
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
)

//记录日志，这个动作的内容只记录一次
func LogActionOnce(action string, v ...interface{}) {
	cache := orm.GetMCache()
	lStr  := "spider@"+Md5Str(action)
	//缓存中已经存在，则不记录
	if isExits := cache.Get(lStr); isExits != nil {
		return
	}
	//相同的内容如果已经记录过了，直接跳过
	cache.Set(lStr, true, 86400)
	log.Write(-1, action)
	log.Write(-1, v...)
}