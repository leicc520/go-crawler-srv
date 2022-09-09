package proxy

import (
	"fmt"
	"github.com/leicc520/go-orm/log"
	"net/http"
	"os"
	"testing"
)

func TestNamev2(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	//link := "http://10.71.32.68:22336/_/www.amazon.com/s?me=AVWPR46RBFNHJ&language=en_US&page=1"
	link := "https://www.amazon.com/b?node=3564021011&ref=sr_nr_n_11"
	//link  = "http://test.loc/index.php?aaa=xxx"
	client := NewHttpRequest()
	client.Proxy("http://prorac2020:c02396-2f7787-27beef-1b857f-1ed432@global.rotating.proxyrack.net:9000")
	client.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	client.AddHeader("Accept-Encoding", "gzip, deflate, br")
	client.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	//client.AddHeader("cache-control", "max-age=0")
	client.AddHeader("Connection", "keep-alive")
	client.AddHeader("Upgrade-Insecure-Requests", "1")
	//client.AddHeader("cookie", "session-id-time=2082787201l; session-id=147-4283500-0346022; ubid-main=135-6328515-0925059; i18n-prefs=USD; lc-main=en_US; sp-cdn=\"L5Z9:SG\"; csm-hit=tb:2HTZ3C7KB5HB3VKMK8WZ+s-15SJPXBHQ1YXZJCV00HW|1661775037176&t:1661775037231&adb:adblk_no; session-token=bkJce5+z2SwWoePRolC2yvevKgOSTXc3B5FdvQXFhgYCG5ZzAbuOYb/mjB7/t3bz0EYFI7mxlsco6euyxj017z1knS3Rfkv8mqYqbeWEaf50gcvLIaMY9hwqDUx9vbc2Uzxm3KFDSih8z6zJN7/PtvilaouV9/sY1WT0F8jFs6/WbX43Pl4Nh0awcryQEMkWQCcLtg84vnEw8myTGDNvEQ4OQe7MDb7K")
	client.AddHeader("TE","trailers")
	client.AddHeader("Sec-Fetch-Dest","document")
	client.AddHeader("Sec-Fetch-Mode","navigate")
	client.AddHeader("Sec-Fetch-Site","none")
	client.AddHeader("Sec-Fetch-User","?1")
	client.AddHeader("User-Agent", UserAgent())
	//获取店铺信息的时候 需要使用session-id
	for i := 0; i < 2; i++ {
		result, err := client.Request(link, nil, "GET")
		fmt.Println(client.Header())
		sessid :=  client.GetJarCookie("https://www.amazon.com/", "session-id")
		fmt.Println(i, err, sessid, result, client.Header(), "=======")
		if err == nil {
			break
		}
		//client.ResetAgent()
		client.AddHeader("cookie", "session-id="+sessid)
	}
}

func TestDemo(t *testing.T) {
	client := NewHttpRequest()
	//client.SetTlsV2("./amazon.pem")
	//client.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	//client.AddHeader("Accept-Encoding", "gzip, deflate, br")
	//client.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	//client.AddHeader("cache-control", "max-age=0")
	client.AddHeader("Connection", "keep-alive")
	client.AddHeader("Upgrade-Insecure-Requests", "1")
	//client.AddHeader("cookie", "session-id-time=2082787201l; session-id=147-4283500-0346022; ubid-main=135-6328515-0925059; i18n-prefs=USD; lc-main=en_US; sp-cdn=\"L5Z9:SG\"; csm-hit=tb:2HTZ3C7KB5HB3VKMK8WZ+s-15SJPXBHQ1YXZJCV00HW|1661775037176&t:1661775037231&adb:adblk_no; session-token=bkJce5+z2SwWoePRolC2yvevKgOSTXc3B5FdvQXFhgYCG5ZzAbuOYb/mjB7/t3bz0EYFI7mxlsco6euyxj017z1knS3Rfkv8mqYqbeWEaf50gcvLIaMY9hwqDUx9vbc2Uzxm3KFDSih8z6zJN7/PtvilaouV9/sY1WT0F8jFs6/WbX43Pl4Nh0awcryQEMkWQCcLtg84vnEw8myTGDNvEQ4OQe7MDb7K")
	//client.AddHeader("TE","trailers")
	//client.AddHeader("Sec-Fetch-Dest","document")
	//client.AddHeader("Sec-Fetch-Mode","navigate")
	//client.AddHeader("Sec-Fetch-Site","none")
	//client.AddHeader("Sec-Fetch-User","?1")
	client.AddHeader("anti-csrftoken-a2z", "gG3zHahEF/VlqD4rfAuimGip5l88/MTuqIB2PhEAAAAMAAAAAGMXMTxyYXcAAAAA;hFYEKl/enMkJ2O2jIwl6nai4QvbLJshTBFP36EEX/wohAAAAAGMXMTwAAAAB")
	//client.AddHeader("User-Agent", RandUserAgent())
	client.Proxy("http://147.185.238.169:50000")
	body := []byte("locationType=LOCATION_INPUT&zipCode=10004&storeContext=generic&deviceType=web&pageType=Search&actionSource=glow&almBrandId=undefined")
	//client.AddHeader("referer", "https://www.amazon.com/s?k=gaming+chair&refresh=6&ref=glow_cls")
	//client.AddHeader("content-length", strconv.FormatInt(int64(len(body)), 10))
	//client.AddHeader("x-requested-with", "XMLHttpRequest")
	client.AddHeader("content-type", "application/x-www-form-urlencoded")

	//client.SetCookie("https://www.amazon.com/")
	client.AddHeader("cookie", "session-id=134-0436097-5771919; session-id-time=2082787201l; i18n-prefs=USD; ubid-main=134-4940307-4482913; lc-main=en_US; skin=noskin; csm-hit=tb:C64YQRWXBK8AWA2SPY5J+s-V5HMW76VXFHH6ZYQ76SW|1662464815301&t:1662464815301&adb:adblk_no; session-token=y4QfsRoE3uif4LKHbpR7Sin+3jkoSHt1Ay/q538jFeCKeFcxwukNDa32HKCcYYNbQHzyj1uGQaxOib1B76NrAfGfIt0Sy6uzcALeBHvuTEc+5ZmsJ8y7SYS8NBQhauWZBa20fcahSpuUl1jHWGnc9mvcyOjEXj1MVcHB4vf9KpLwn5DpQ9xUQNla9Icpa+jjlr/suWTnxwNnnb2SFHg61vIc144AezxG")

	result, err := client.Request("https://www.amazon.com/gp/delivery/ajax/address-change.html", body, http.MethodPost)
	fmt.Println(result, string(body), err)
	fmt.Printf("%+v", client.Header())

}

func TestName(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	type TaskSt struct {
		Url string
		Method string
		Body string
		Header map[string]string
	}
	tasks := []TaskSt{
		{Url: "https://www.amazon.com/", Method: http.MethodGet},
		{Url: "https://www.amazon.com/gp/delivery/ajax/address-change.html",
			Body: "locationType=LOCATION_INPUT&zipCode=10003&storeContext=generic&deviceType=web&pageType=Search&actionSource=glow&almBrandId=undefined",
			Method: http.MethodPost, Header: map[string]string{"referer":"https://www.amazon.com/"}},
		{Url: "https://www.amazon.com/s/query?k=gaming%20chair&page=1&ref=glow_cls&refresh=4", Body: "{\"customer-action\":\"query\"}",
			Method: http.MethodPost, Header: map[string]string{"referer":"https://www.amazon.com/"}},
	}

	client := NewHttpRequest()
	//client.SetTlsV2("./amazon.pem")
	client.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	client.AddHeader("Accept-Encoding", "gzip, deflate, br")
	client.AddHeader("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	//client.AddHeader("cache-control", "max-age=0")
	client.AddHeader("Connection", "keep-alive")
	client.AddHeader("Upgrade-Insecure-Requests", "1")
	//client.AddHeader("cookie", "session-id-time=2082787201l; session-id=147-4283500-0346022; ubid-main=135-6328515-0925059; i18n-prefs=USD; lc-main=en_US; sp-cdn=\"L5Z9:SG\"; csm-hit=tb:2HTZ3C7KB5HB3VKMK8WZ+s-15SJPXBHQ1YXZJCV00HW|1661775037176&t:1661775037231&adb:adblk_no; session-token=bkJce5+z2SwWoePRolC2yvevKgOSTXc3B5FdvQXFhgYCG5ZzAbuOYb/mjB7/t3bz0EYFI7mxlsco6euyxj017z1knS3Rfkv8mqYqbeWEaf50gcvLIaMY9hwqDUx9vbc2Uzxm3KFDSih8z6zJN7/PtvilaouV9/sY1WT0F8jFs6/WbX43Pl4Nh0awcryQEMkWQCcLtg84vnEw8myTGDNvEQ4OQe7MDb7K")
	client.AddHeader("TE","trailers")
	client.AddHeader("Sec-Fetch-Dest","document")
	client.AddHeader("Sec-Fetch-Mode","navigate")
	client.AddHeader("Sec-Fetch-Site","none")
	client.AddHeader("Sec-Fetch-User","?1")
	client.AddHeader("User-Agent", UserAgent())
	client.Proxy("http://147.185.238.169:50000")
	for idx, task := range tasks {
		var body []byte = nil
		if len(task.Body) > 0 {
			body = []byte(task.Body)
		}
		result, err := client.Request(task.Url, body, task.Method)
		if err != nil {
			fmt.Println(task, err)
			break
		}
		fmt.Println("step====", idx)
		if idx == 2 {
			fmt.Println(result)
		}
	}

	/*
	//获取店铺信息的时候 需要使用session-id
	body := []byte("{\"customer-action\":\"query\"}")
	for i := 0; i < 2; i++ {
		result, err := client.Request(link, body, "POST")
		//fmt.Println(client.Header())
		sessid :=  client.GetJarCookie("https://www.amazon.com/", "session-id")
		fmt.Println(i, err, sessid, result, "=======")
		if err == nil {
			break
		}
		//client.ResetAgent()
		client.AddHeader("cookie", "session-id="+sessid)
	}*/
}


func TestNameV2(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	//link := "http://10.71.32.68:22336/_/www.amazon.com/sp?seller=A2VTI8ZSUTN12G&language=en_US"
	//link := "https://www.amazon.com/sp?seller=A2VTI8ZSUTN12G&language=en_US"
	//link := "https://www.amazon.com/dp/B08C74Y5L2?language=en_US&psc=1&th=1"
	link := "http://10.71.32.68:22336/_/www.amazon.com.mx/sp?seller=AC7MYPIY9FC19&language="
	//link := "https://www.amazon.ca/sp?seller=A3NSIIMM24A1EP&language=en_US"  https://
	//link    = "http://10.100.72.102:80/index.php?aaa=xxx"
	client := NewHttpRequest()
	//client.SetTlsV2("./amazon.pem")
	client.AddHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	client.AddHeader("accept-encoding", "gzip, deflate, br")
	client.AddHeader("accept-language", "en")
	//client.AddHeader("cache-control", "max-age=0")
	client.AddHeader("connection", "keep-alive")
	client.AddHeader("upgrade-insecure-requests", "1")
	//client.AddHeader("cookie", "session-id-time=2082787201l; session-id=147-4283500-0346022; ubid-main=135-6328515-0925059; i18n-prefs=USD; lc-main=en_US; sp-cdn=\"L5Z9:SG\"; csm-hit=tb:2HTZ3C7KB5HB3VKMK8WZ+s-15SJPXBHQ1YXZJCV00HW|1661775037176&t:1661775037231&adb:adblk_no; session-token=bkJce5+z2SwWoePRolC2yvevKgOSTXc3B5FdvQXFhgYCG5ZzAbuOYb/mjB7/t3bz0EYFI7mxlsco6euyxj017z1knS3Rfkv8mqYqbeWEaf50gcvLIaMY9hwqDUx9vbc2Uzxm3KFDSih8z6zJN7/PtvilaouV9/sY1WT0F8jFs6/WbX43Pl4Nh0awcryQEMkWQCcLtg84vnEw8myTGDNvEQ4OQe7MDb7K")
	client.AddHeader("TE","trailers")
	client.AddHeader("sec-fetch-dest","document")
	client.AddHeader("sec-fetch-mode","navigate")
	client.AddHeader("sec-fetch-site","none")
	client.AddHeader("sec-fetch-user","?1")


	client.AddHeader("user-agent", UserAgent())

	client.AddHeader("X-Zbox-Auth-Token", os.Getenv("CrawlerProxyTokenPool"))

	//获取店铺信息的时候 需要使用session-id
	for i := 0; i < 2; i++ {
		result, err := client.Request(link, nil, "GET")
		fmt.Println(client.Header())
		sessid :=  client.GetJarCookie("https://www.amazon.com/", "session-id")
		fmt.Println(i, err, sessid, result, client.GetResponse(), "=======")
		break
		if err == nil {
			break
		}
		//client.ResetAgent()
		client.AddHeader("cookie", "session-id="+sessid)
	}
}
