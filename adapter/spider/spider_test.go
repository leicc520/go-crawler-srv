package spider

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/lib/parser"
	"github.com/leicc520/go-orm/cache"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func init() {
	os.Chdir("../../")
	dir, err := os.Getwd()
	fmt.Println(dir, err)
	lib.InitConfig()
	cConfig:= map[string]interface{}{"dir":"./cachedir/webcache", "dept":2}
	cCache := cache.Factory("file", cConfig)
	parser.Inject(cCache, nil)
}

func TestSeller(t *testing.T) {
	tt := parser.TemplateSt{Request: &parser.BaseRequest{}}
	err:= tt.LoadFile("./config/template/amazon-seller.yml")
	fmt.Println(err)
	fmt.Printf("%+v %+v", tt.Request, tt)
	link := "https://www.amazon.com/sp?ie=UTF8&seller=ADJ8GJXNTQ637&asin=B004809ITC&ref_=dp_merchant_link&isAmazonFulfilled=1"
	result, err := tt.Request.Do(link)
	fmt.Println(err)
	item, err := parser.NewCompiler(result).Parse(tt.DataFields)
	fmt.Printf("%+v %+v", item, err)
}