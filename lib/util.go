package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/leicc520/go-orm/log"
)

//读取文件内容数据信息
func ReadFile(file string) string {
	stream, err := ioutil.ReadFile(file)
	if err != nil {
		log.Write(log.ERROR, "file read error", err)
	}
	return string(stream)
}

//清理html 样式表和js代码
func HTMLClean(htmlStr string) string {
	htmlStr = regexp.MustCompile(`(?s)<style.*?>.*?</style>`).ReplaceAllString(htmlStr, "")
	htmlStr = regexp.MustCompile(`(?s)<noscript.*?>.*?</noscript>`).ReplaceAllString(htmlStr, "")
	htmlStr = regexp.MustCompile(`(?s)<script.*?>.*?</script>`).ReplaceAllString(htmlStr, "")
	return htmlStr
}

//过滤html标签处理逻辑
func StripTags(htmlStr string) string {
	htmlStr = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(htmlStr, "")
	htmlStr = NormalizeSpace(htmlStr)
	return htmlStr
}

//将空格合并，过滤前后空格
func NormalizeSpace(s string) string {
	return strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(s, " "))
}

//字符串截取字段逻辑
func CutStr(str string, length int, suffix string) string {
	s := []rune(str)
	total := len(s)
	if total <= length {
		return str
	}
	if length < 0 {
		length = total
	}
	result := string(s[0:length])+suffix
	return result
}

//获取字符串md5 hash值
func Md5Str(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//格式化json数据资料信息
func PrettyJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	return out.String()
}
