package lib

import (
	"errors"
	"github.com/leicc520/go-orm/log"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
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

//字符串解析获取int类型的字符
func ParseInt(s string) (n int, err error) {
	s = regexp.MustCompile(`\d+`).FindString(s)
	if s == "" {
		err = errors.New("获取Int解析异常")
		return
	}
	return strconv.Atoi(s)
}

//解析获取所有的整型数据
func ParseIntAll(s string) (numbers []int, err error) {
	aStr := regexp.MustCompile(`\d+`).FindAllString(s, -1)
	if aStr == nil || len(aStr) < 1 {
		err = errors.New("获取Int切片不存在")
	}
	tmpInt := 0
	for _, str := range aStr {
		tmpInt, err = strconv.Atoi(str)
		if err != nil {
			return
		}
		numbers = append(numbers, tmpInt)
	}
	return
}

//提取解析浮点数信息
func ParseFloat64(s string) (n float64, err error) {
	floatStr := regexp.MustCompile(`\d+\.\d+|\d+`).FindString(s)
	if floatStr == "" {
		err = errors.New("获取Float解析异常")
		return
	}
	return strconv.ParseFloat(floatStr, 64)
}

//提取解析所有的浮点数
func ParseFloat64All(s string) (numbers []float64, err error) {
	aStr := regexp.MustCompile(`\d+\.\d+|\d+`).FindAllString(s, -1)
	var number float64
	for _, str := range aStr {
		number, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return
		}
		numbers = append(numbers, number)
	}
	return
}

//将空格合并，过滤前后空格
func NormalizeSpace(s string) string {
	return strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(s, " "))
}

//获取页数信息
func CeilDivPage(total, numPage int) int {
	return int(math.Ceil(float64(total) / float64(numPage)))
}
