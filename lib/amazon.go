package lib

import (
	"errors"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/leicc520/go-gin-http/micro"
)

var (
	usZipCode  = [][][]int{
		{{20001, 20013},{20015, 20020},{20035,20042}},    //washington
		{{10001,10014}, {10016,10020}, {10035,10041}},   //newyork
		{{60601,60626}},  //chicago
	}
	AmazonConfig  = struct{
		SearchSiteSupported string `yaml:"search_site_supported"`
		Amazon AmazonSiteSt `yaml:"amazon"`
	}{}
)

type AmazonSiteItemSt struct {
	CName 	string  `yaml:"cname"`
	UName 	string  `yaml:"uname"`
	Code  	string  `yaml:"code"`
	Suffix 	string 	`yaml:"suffix"`
	Lang  	string  `yaml:"lang"`
	Cny   	string  `yaml:"cny"`
	ZipCode string 	`yaml:"zip_code"`
	City    string 	`yaml:"city"`
}

type AmazonSiteSt []AmazonSiteItemSt

//加载亚马逊的配置信息
func amazonConfigLoad(configFile string) error {
	_, err := micro.LoadFile(configFile, &AmazonConfig)
	return err
}

//获取地址数据资料信息
func (s AmazonSiteSt) GetSiteByUrl(src string) (*AmazonSiteItemSt, error) {
	uri, err := url.Parse(src)
	if err != nil {
		return nil, err
	}
	for _, item := range s {
		if strings.HasSuffix(uri.Host, item.Suffix) {
			return &item, nil
		}
	}
	return nil, errors.New(src+" 站点未作适配,无法抓取")
}

//获取这个站点地区的物流验证码和地址
func (s *AmazonSiteItemSt) GenZipCode() (string, string) {
	rand.Seed(time.Now().UnixNano())
	switch s.Code {
	case "US":
		return s.usZipCode(), ""
	}
	zipCode, zipCity := "", ""
	idx := rand.Int()
	if len(s.ZipCode) > 0 {
		zipCodes := strings.Split(s.ZipCode, ",")
		idx       = idx%len(zipCodes)
		zipCode   = zipCodes[idx]
	}
	if len(s.City) > 0 {
		zipCitys := strings.Split(s.City,",")
		if idx < len(zipCitys) && len(zipCode) > 0 {
			zipCity = zipCitys[idx]
		}
	}
	return zipCode, zipCity
}

//获取美国的主要国家编码
func (s *AmazonSiteItemSt) usZipCode() string {
	fIdx := rand.Intn(len(usZipCode))
	sIdx := rand.Intn(len(usZipCode[fIdx]))
	item := usZipCode[fIdx][sIdx]
	return s.randRange(item[0], item[1])
}

//获取数值范围数据信息
func (s *AmazonSiteItemSt) randRange(min, max int) string {
	rsize := rand.Int() % (max - min + 1) + min
	return strconv.FormatInt(int64(rsize), 10)
}