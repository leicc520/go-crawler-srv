package proxy

import (
	"fmt"
	"github.com/leicc520/go-orm/log"
	"os"
	"testing"
)

func TestNamev2(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	//link := "http://10.71.32.68:22336/_/www.amazon.com/s?me=AVWPR46RBFNHJ&language=en_US&page=1"
	link := "https://www.amazon.com/product-reviews/B09VFZ6PKT?sortBy=recent&pageNumber=2&language=en_US&filterByStar=all_stars&pageSize=20"
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
	client.AddHeader("User-Agent", RandUserAgent())
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

func TestName(t *testing.T) {
	log.Write(log.DEBUG, "11111111111111111")
	//link := "http://10.71.32.68:22336/_/www.amazon.com/s?me=AVWPR46RBFNHJ&language=en_US&page=1"
	link := "https://www.amazon.com/sp?seller=A2VTI8ZSUTN12G&language=en_US"
	//link  = "http://test.loc/index.php?aaa=xxx"
	client := NewHttpRequest()
	client.SetTlsV2("./amazon.pem")
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


	client.AddHeader("User-Agent", RandUserAgent())

	//client.AddHeader("X-Zbox-Auth-Token", os.Getenv("CrawlerProxyTokenPool"))
	//client.AddHeader("origin","https://www.amazon.com")
	//client.AddHeader("authority", "www.amazon.com")
	//client.AddHeader("referer", "https://www.amazon.com/")
	//client.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:104.0) Gecko/20100101 Firefox/104.0")

	//result, err := client.Request("https://www.amazon.com/", nil, "GET")
	//fmt.Println(err, result, client.GetResponse(), client.Header(), "=======")
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


	client.AddHeader("user-agent", RandUserAgent())

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
