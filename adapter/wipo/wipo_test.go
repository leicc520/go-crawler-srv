package wipo

import (
	"encoding/json"
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/proxy"
	"github.com/leicc520/go-crawler-srv/lib/proxy/channal"
	"github.com/leicc520/go-crawler-srv/plugins"
	"testing"
	"time"

	_ "github.com/lib/pq"
)
import "github.com/Lazarus/lz-string-go"

func TestEncrypt(t *testing.T) {
	str := "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCAKoDKIANCAGYCWJAQgEoCCAcgCIgC+AujVo5I+Oo0ghyAFTbSK1EERJsAwtMUoA9iQDyLPv140kWgO4woANgAMNGAWQFrdughRwCl0SDY8jNAQAnmDEkgBGSAgAdgAmigA2CCRw0YpYEiAATAC0AKJgtABStADsACwAGhy6ALIJ5EEAzCzSHAAeZrQsNu157eFgAFpWBFnlugCuAJoAXuUA5nAsRXkAvOmTcFtQAIw0APoku7sAnE1ZfEAAA=="
	res, err := LZString.Decompress(str, "")
	fmt.Println(res, err)
	
	//qi
	//:
	//"7-auLhQbGY9nkJyJ6C2Mu4uFsVeeLL9rvw3bxCwJ5qQ/E="
	//qk
	//:
	//"Pg36ac6Yoe0Cfav1rh5rMejV9UY9sSw7y1owNaFzoT8="
	str = "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCLgEwAMZZAtBQJx30AqFFkbHFAWgATMB5XpWpMmzMgGZIAVnqz63ALogANCABmASxIAlACJqQAEwIkRtBkxABfJTfVIA9gHcYUAGwV1MAsjOQAIyUDiAEAJ5gxJAgAEZICAB2xkYANggkcIlGWDoxABw0AAoA5pIeCCgeAJpOcBQAwhoIAG6BSGgySACycABWAGr0AKrV9DAAyi4A7OGBrgByCABiAF5OzPkAvDkArnD7UIHqAPokgYH0kmS2QA"
	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)

	str = "N4IgDiBcoGYDZRAQQCIBoBCAlJA5dAougBIDyAMmmeQMJoCS69AsgOJrmmulq52k8s6AMqkAqlhoE0wgCpJZY4WgDOAYwD2AJwCmDdClI0QAXzQgVUUCp0AXKyDWXIjgBY61Aa0w58M8ZLScgpKqpq6aAJUFNG0 mioHFw8fAxsaEJoRKZmILYAnmA6iABGWgCGAHYAJiDmcOWIOpV1IACOAJaIAIwAtOXVAAwArrZi1Z4AXjAAnAAaAO4ArGqkAIq4CzTEC QAYtVrKhhtAFKuw8NrpG2yABw0YgDmzJ4AvK1twzrfUN3mAH0et0ZgBmABMpiAA"

	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)

	str = "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCAKoDKIANCAGYCWJA8kyAL4C6bNtANiQEEAIlQBCAJQEA5EQFERACSYAZKkuUBhKgEkR2gLIBxKsqaGmVKVpZVxI8k1LiNsquQAqA9xSowUAeyQ4HREhJg12GhhMEHgCGJRoyBB0OBQAazFJGTdHZ1cPLx8/QOCbdTUVLV0qYRMzCysdI1s5IXZuEAIATzBiZIAjJAQAOwATahBeBBI4EcmsRmSARgBaBDGABgBXAlIx9IAvWgBOAA0AdwBWFCYARSkLjQUL5QAxMbuYUSwAKTRtts7kwsO4ABwaUgAc306QAvAttnAkVBljQAPokZbLE4AZgATOwgAAA=="
	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)
	//8-Pg36ac6Yoe0Cfav1rh5rMejV9UY9sSw7y1owNaFzoT8=
}

func TestQk(t *testing.T) {
	str := `var w = $(window).width();
var qk = "rx0aodmPMMlXiR/ABFsCBmatZsCMvsgv7XRUXKfGaQI";

// if(!((w == 790 || w == 800) && (h == 600 || h == 590)))
qk = "U+dA1Ga0wy825xWj9w5cJixmSwAZPUXwMq9k5KdjR0Q=";`
	wepi := (&WipoSt{}).parseQk2QiString(str)
	fmt.Println(wepi)
}

func TestCookie(t *testing.T) {
	str := "demotest"
	fmt.Println(wipoVisitorUunId(str))
}

func TestDate(t *testing.T) {
	ss := time.Now()
	body, err := json.Marshal(ss)
	
	fmt.Println(string(body), err)
	
	bb := time.Time{}
	err = json.Unmarshal(body, &bb)
	fmt.Println(err, bb)
}

func TestWipo(t *testing.T) {
	lib.InitRedis("redis://:@127.0.0.1:6379/1")



	//初始化数据资料信息
	//{Url: "", Proxy: "easy-go-http", Status: 1, IFGet: &channal.EasyGoSt{}}
	//{Url: "", Proxy: "sky-go-http", Status: 1, IFGet: &channal.SkyStartGoSt{}}
	proxyHost := []proxy.ProxySt{{Url: "", Proxy: "easy-go-http", Status: 1, IFGet: &channal.EasyGoSt{}}}
	proxy.Init(proxyHost, lib.Redis)
	ss := &WipoSt{dpc:&plugins.ChromeDpSt{HeadLess: false}, Country: "US"}
	ss.Run("2022-01-01", "2022-09-14")
}
