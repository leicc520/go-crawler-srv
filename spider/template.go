package spider

import (
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-crawler-srv/spider/parse"
	"github.com/leicc520/go-orm/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

/*********************************************************************
 	配置模板参数数据资料信息，只有user-agent是随机的，其他走配置
 */

type TemplateSt struct{
	Method    string           `json:"method"   yaml:"method"`
	Params    string           `json:"params"   yaml:"params"`
	Headers   proxy.HeaderSt   `json:"headers"  yaml:"headers"`
	Elements []parse.ElementSt `json:"elements" yaml:"elements"`
}

//加载配置数据资料信息
func (s *TemplateSt) LoadFile(confFile string)  error {
	if file, err:=os.Stat(confFile); err != nil || !file.Mode().IsRegular() {
		log.Write(log.ERROR, "load Template Config File Failed: ", err)
		return err
	}
	data, _ := ioutil.ReadFile(confFile)
	//把yaml形式的字符串解析成struct类型 先子类初始化
	if err := yaml.Unmarshal(data, s); err != nil {
		log.Write(log.ERROR, "load Template Config Parse Failed: ", err)
		return err
	}
	return nil
}

//抓取网页数据处理逻辑
func (s *TemplateSt) Crawling(url string) string {
	client := proxy.NewHttpRequest().SetHeader(s.Headers.ASMap())
	method := strings.ToUpper(s.Method)
	doc, err := client.Request(url, []byte(s.Params), method)
	if err != nil {
		return ""
	}
	return doc
}

//配置模板数据资料信息
type TemplatesSt map[string]TemplateSt

//获取模板名称数据资料信息
func (s TemplatesSt) getTemplate(templateName string) []parse.ElementSt {
	if s == nil || len(s) < 1 {
		return nil
	}
	if tmpl, ok := s[templateName]; ok {
		return tmpl.Elements
	}
	return nil
}