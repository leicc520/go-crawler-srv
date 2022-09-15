package plugins

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync/atomic"

	"github.com/chromedp/chromedp"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	logsf "github.com/leicc520/go-orm/log"
)

type ActionCb func(string, context.Context) (string, error)

type ChromeDpSt struct {
	statistic uint64 	`yaml:"-"`  //统计数值
	Agent   string 		`yaml:"-"`   //浏览器
	IsDebug  bool 		`yaml:"is_debug"`
	HeadLess bool 		`yaml:"head_less"`
	ProxyUrl string 	`yaml:"proxy_url"`
	DevtoolsWs []string `yaml:"devtools_ws"`
}

//获取原创的headless
func (s *ChromeDpSt) getDevToolsWs() string {
	size := uint64(len(s.DevtoolsWs))
	n := atomic.AddUint64(&s.statistic, 1)
	if size > 0 {
		return s.DevtoolsWs[n%size]
	}
	return ""
}

//初始化数据资料信息 配置选项
func (s *ChromeDpSt) options() []chromedp.ExecAllocatorOption {
	s.Agent    = proxy.UserAgent()
	options   := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", s.HeadLess), // debug使用
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		chromedp.Flag("blink-settings", "imagesEnabled=false"), // 禁用图片加载
		chromedp.WindowSize(1920, 1080),
		chromedp.UserAgent(s.Agent),
	}
	if len(s.ProxyUrl) > 0 { //设置开启代理的情况处理逻辑
		options = append(options, chromedp.ProxyServer(s.ProxyUrl))
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	return options
}

//执行发起一个网络业务请求 @sel 需要返回元素的选择器 默认html
func (s *ChromeDpSt) Run(url string, taskCb ActionCb) (htmlDoc string, err error) {
	defer func() { //异常的捕获处理逻辑
		if pErr := recover(); pErr != nil {
			err  = errors.New(fmt.Sprintf("%+v", pErr))
			logsf.Write(logsf.ERROR, err, "异常处理逻辑...")
		}
	}()
	var aCtx context.Context = nil
	var aCancel context.CancelFunc = nil
	devToolsWs := s.getDevToolsWs()
	if len(devToolsWs) > 0 {//使用远程代理请求的情况
		aCtx, aCancel = chromedp.NewRemoteAllocator(context.Background(), devToolsWs)
	} else {
		aCtx, aCancel = chromedp.NewExecAllocator(context.Background(), s.options()...)
	}
	defer aCancel()
	logDebug := func (format string, v ...interface{}){}
	if s.IsDebug {//调式模式的情况
		logDebug = log.Printf
	}
	ctx, rCancel  := chromedp.NewContext(aCtx, chromedp.WithDebugf(logDebug))
	defer rCancel()
	htmlDoc, err = taskCb(url, ctx)
	if err != nil { //处理执行结果处理逻辑
		logsf.Write(logsf.ERROR, "chromedp.Run failed", err)
	}
	return
}