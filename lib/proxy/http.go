package proxy

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/leicc520/go-orm/log"
)

const CONTENT_TYPE = "content-type"

type HttpSt struct {
	sp           *http.Response
	query        url.Values
	isRedirect   bool
	timeout      time.Duration
	cookieJar    *cookiejar.Jar
	tlsTransport http.RoundTripper
	header       map[string]string
}

func CancelRedirect(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}

func NewHttpRequest() *HttpSt {
	cookieJar, _ := cookiejar.New(nil)
	return &HttpSt{query: url.Values{}, isRedirect: true, header: nil,
		timeout: 120 * time.Second, cookieJar: cookieJar}
}

//设置请求的header业务数据信息
func (s *HttpSt) SetTimeout(timeout int) *HttpSt {
	s.timeout = time.Duration(timeout) * time.Second
	return s
}

//设置请求的header业务数据信息
func (s *HttpSt) SetHeader(header map[string]string) *HttpSt {
	if s.header != nil { //数据不为空的情况
		for key, val := range header {
			s.header[key] = val
		}
	} else {
		s.header = header
	}
	return s
}

//初始化请求数据头部信息
func (s *HttpSt) initHeader(req *http.Request) *HttpSt {
	baseUrl  := fmt.Sprintf("%s://%s", req.URL.Scheme, req.URL.Host)
	s.AddHeader("origin", baseUrl)
	if _, ok := s.header["accept"]; !ok { //不存在的话设置一下
		s.AddHeader("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	}
	if _, ok := s.header["accept-encoding"]; !ok { //不存在的话设置一下
		s.AddHeader("accept-encoding", "gzip, deflate, br")
	}
	if _, ok := s.header["upgrade-insecure-requests"]; !ok { //不存在的话设置一下
		s.AddHeader("upgrade-insecure-requests", "1")
	}
	if _, ok := s.header["sec-fetch-dest"]; !ok { //不存在的话设置一下
		s.AddHeader("sec-fetch-dest", "document")
	}
	if _, ok := s.header["sec-fetch-mode"]; !ok { //不存在的话设置一下
		s.AddHeader("sec-fetch-mode", "navigate")
	}
	if _, ok := s.header["sec-fetch-site"]; !ok { //不存在的话设置一下
		s.AddHeader("sec-fetch-site", "none")
	}
	if _, ok := s.header["sec-fetch-user"]; !ok { //不存在的话设置一下
		s.AddHeader("sec-fetch-user", "?1")
	}
	if _, ok := s.header["accept-language"]; !ok { //不存在的话设置一下
		s.AddHeader("accept-language", "*")
	}
	if _, ok := s.header["referer"]; !ok { //不存在的话设置一下
		s.AddHeader("referer", baseUrl+"/")
	}
	if _, ok := s.header["user-agent"]; !ok { //不存在的话设置一下
		s.AddHeader("user-agent", RandUserAgent())
	}
	return s
}

//重置浏览器agent数据
func (s *HttpSt) ResetAgent() *HttpSt {
	s.AddHeader("user-agent", RandUserAgent())
	return s
}

//设置是否重定向处理逻辑,默认true
func (s *HttpSt) IsRedirect(is bool) *HttpSt {
	s.isRedirect = is
	return s
}

//获取指定的cookie信息
func (s *HttpSt) GetJarCookie(link, name string) string {
	u, _ := url.Parse(link)
	cookies := s.cookieJar.Cookies(u)
	for _, item := range cookies {
		if item.Name == name {
			return item.Value
		}
	}
	return ""
}

//返回数据记录信息
func (s *HttpSt) GetResponse() *http.Response {
	return s.sp
}

//设置请求的header业务数据信息
func (s *HttpSt) AddHeader(key, val string) *HttpSt {
	if s.header == nil {
		s.header = map[string]string{}
	}
	s.header[key] = val
	return s
}

//设置发起json的业务请求json,xml,default
func (s *HttpSt) SetContentType(typeStr string) *HttpSt {
	if s.header == nil {
		s.header = map[string]string{}
	}
	switch strings.ToLower(typeStr) {
	case "json":
		s.header[CONTENT_TYPE] = "application/json; charset=utf-8"
	case "xml":
		s.header[CONTENT_TYPE] = "application/xml; charset=utf-8"
	default:
		s.header[CONTENT_TYPE] = "application/x-www-form-urlencoded"
	}
	return s
}

//添加设置查询语句
func (s *HttpSt) Set(name, value string) *HttpSt {
	s.query.Set(name, value)
	return s
}

//获取设置的http请求header数据
func (s *HttpSt) Header() map[string]string {
	return s.header
}

//获取查询的语句数据
func (s *HttpSt) Query() string {
	return s.query.Encode()
}

//重置请求的参数数据信息
func (s *HttpSt) Reset() *HttpSt {
	s.query = url.Values{}
	s.header = nil
	return s
}

//重置请求的参数数据信息
func (s *HttpSt) SetTls(keySsl, pemSsl string) *HttpSt {
	c, _ := tls.X509KeyPair([]byte(pemSsl), []byte(keySsl))
	cfg := &tls.Config{
		Certificates: []tls.Certificate{c},
	}
	s.tlsTransport = http.RoundTripper(&http.Transport{
		TLSClientConfig: cfg,
	})
	return s
}

//重置请求的参数数据信息
func (s *HttpSt) SetTlsV2(pemSsl string) *HttpSt {
	caCert, err := ioutil.ReadFile(pemSsl)
	if err != nil {
		log.Write(log.ERROR, "SetTlsV2", err)
		return s
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
		InsecureSkipVerify: true,
	}
	s.tlsTransport = http.RoundTripper(&http2.Transport{
		TLSClientConfig: tlsConfig,
	})
	return s
}

//设置启动http代理发起业务请求
func (s *HttpSt) Proxy(proxyUrl string) *HttpSt {
	uri, err := url.Parse(proxyUrl)
	if err != nil {
		log.Write(log.ERROR, "set proxy error", proxyUrl, err)
		return s
	}
	t := &http.Transport{TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	}, TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)}
	t.Proxy = http.ProxyURL(uri)
	s.tlsTransport = http.RoundTripper(t)
	return s
}

//上传文件处理逻辑 封装成byte
func (s *HttpSt) UpFile(param map[string]string, paramName, path, fileName string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if fileName == "" {
		fileName = filepath.Base(path)
	}
	fp, err := writer.CreateFormFile(paramName, fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fp, file)
	for key, val := range param {
		_ = writer.WriteField(key, val)
	}
	s.SetHeader(map[string]string{"content-type": writer.FormDataContentType()})
	if err = writer.Close(); err != nil {
		return nil, err
	}
	return body.Bytes(), nil
}

//请求下载文件数据信息
func (s *HttpSt) DownLoad(url, filePath string) (string, error) {
	var fp *os.File = nil
	var sp *http.Response = nil
	defer func() { //补货异常的处理逻辑
		if sp != nil && sp.Body != nil {
			sp.Body.Close()
		}
		if r := recover(); r != nil {
			log.Write(log.ERROR, "request url ", url, "error", r)
		}
		if fp != nil {
			fp.Close()
		}
	}()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Write(log.ERROR, url, err)
		return "", err
	}
	if s.header != nil && len(s.header) > 0 {
		for key, val := range s.header {
			req.Header.Set(key, val)
		}
	}
	client := &http.Client{Timeout: s.timeout, Jar: s.cookieJar}
	if s.tlsTransport != nil { //设置加密请求业务逻辑
		client.Transport = s.tlsTransport
	}
	if sp, err = client.Do(req); err != nil || sp == nil {
		log.Write(log.ERROR, url, err)
		return "", err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	io.Copy(f, sp.Body)
	return filePath, nil
}

//发起一个http业务请求
func (s *HttpSt) Request(url string, body []byte, method string) (result string, err error) {
	s.sp = nil
	defer func() { //补货异常的处理逻辑
		if s.sp != nil && s.sp.Body != nil {
			s.sp.Body.Close()
		}
		if r := recover(); r != nil {
			log.Write(log.ERROR, "request url ", url, "error", r)
		}
	}()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Write(log.ERROR, url, err, string(body))
		return
	}
	s.initHeader(req) //初始化附加头信息
	if s.header != nil && len(s.header) > 0 {
		for key, val := range s.header {
			req.Header.Set(key, val)
		}
	}
	fmt.Println(req)
	client := &http.Client{Timeout: s.timeout, Jar: s.cookieJar}
	if s.tlsTransport != nil { //设置加密请求业务逻辑
		client.Transport = s.tlsTransport
	}
	//关闭重定向处理逻辑
	if !s.isRedirect {
		client.CheckRedirect = CancelRedirect
	}
	s.sp, err = client.Do(req)
	if err != nil || s.sp == nil {
		log.Write(log.ERROR, url, err, string(body))
		return
	}
	fmt.Println(s.sp)
	fmt.Println(req)
	if s.sp.StatusCode == 200 {
		return s.readResult() //返回请求回来的数据信息
	} else {
		return "", errors.New(s.sp.Status)
	}
}

//读取解码的数据资料信息
func (s *HttpSt) readResult() (result string, err error) {
	var body []byte
	if s.sp.Header.Get("Content-Encoding") == "gzip" {
		body, err = GZIPDecode(s.sp.Body)
	} else {
		body, err = ioutil.ReadAll(s.sp.Body)
	}
	if err != nil {
		log.Write(log.ERROR, s, err, string(body))
		return
	}
	if strings.Contains(s.sp.Header.Get("Content-Type"), "charset=iso-8859-1") {
		result, err = Decode("iso-8859-1", body)
	} else {
		result = string(body)
	}
	return
}
