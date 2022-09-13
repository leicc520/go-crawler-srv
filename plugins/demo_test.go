package plugins

import (
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"os"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	url := "http://test.loc/"
	// 设置邮编对话框
	cb := func(url string, ctx context.Context) (string, error) {
		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.Evaluate("window.alert = function() { return false;}", nil),
			chromedp.WaitVisible(`//*[@id="demo"]`),
			chromedp.Click(`//*[@id="demo"]`),
		})
		return "", nil
	}

	//Mozilla/5.0 (Windows NT 6.2; Win32; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.150 Safari/537.36
	//Mozilla/5.0 (Windows NT 6.2; Win32; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.150 Safari/537.36
	st, err := (&ChromeDpSt{HeadLess: true}).Run(url, cb)
	fmt.Println(st, err)
}

func TestProxy(t *testing.T) {
	url := "https://google.com.hk"
	// 设置邮编对话框
	cb := func(url string, ctx context.Context) (string, error) {
		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.Evaluate("window.alert = function() { return false;}", nil),
		})
		return "", nil
	}

	proxy := os.Getenv("ZXProxy")
	proxy = "101.36.126.19:6579"
	fmt.Println(proxy)
	//Mozilla/5.0 (Windows NT 6.2; Win32; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.150 Safari/537.36
	//Mozilla/5.0 (Windows NT 6.2; Win32; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.150 Safari/537.36
	st, err := (&ChromeDpSt{HeadLess: true, ProxyUrl: proxy}).Run(url, cb)
	fmt.Println(st, err)
}

func TestWipoCookie(t *testing.T) {
	url   := "https://branddb.wipo.int/branddb/en/#"
	cookieStr := make([]string, 0)
	// 设置邮编对话框
	cb := func(url string, ctx context.Context) (string, error) {
		chromedp.Run(ctx, chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.Sleep(time.Second*10),
			chromedp.ActionFunc(func(ctx context.Context) error {
				cookies, err := network.GetAllCookies().Do(ctx)
				if err != nil {
					return err
				}
				for _, cookie := range cookies {
					cookieStr  = append(cookieStr, cookie.Name+"="+cookie.Value)
				}
				return nil
			}),
		})
		return "", nil
	}
	//Mozilla/5.0 (Windows NT 6.2; Win32; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.150 Safari/537.36
	//Mozilla/5.0 (Windows NT 6.2; Win32; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.150 Safari/537.36
	dpc := &ChromeDpSt{HeadLess: true}
	st, err := dpc.Run(url, cb)
	fmt.Println(st, dpc.Agent, cookieStr, err)
}

func TestGo(t *testing.T) {
	ch := make(chan int)
	go func(ss chan int) {
		select {
		case <-ss:
			fmt.Println("OK-1")
		default:
			for {
				time.Sleep(time.Second)
				fmt.Println("working-1...")
			}
		}
	}(ch)

	go func(ss chan int) {
		select {
		case <-ss:
			fmt.Println("OK-2")
		default:
			for {
				time.Sleep(time.Second)
				fmt.Println("working-2...")
			}
		}
	}(ch)

	time.Sleep(time.Second*3)
	close(ch)
	time.Sleep(time.Second*3)
}

func asyncDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))
	defer func() {
		fmt.Println("====cancel")
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			fmt.Println("working...")
		}
		// 发送HTTP请求
	}(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!!")

	case <-time.After(time.Duration(time.Second * 6)):
		fmt.Println("timeout!!!")

	}
}

func TestContext(t *testing.T) {
	go asyncDemo()
	for {
		time.Sleep(time.Second)
		fmt.Println("working222...")
	}
}
