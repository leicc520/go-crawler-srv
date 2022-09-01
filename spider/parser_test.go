package spider

import (
	"fmt"
	"os"
	"testing"
)

func TestXPathTest(t *testing.T) {
	os.Chdir("../")
	dir, err := os.Getwd()
	fmt.Println(dir, err)

	tt := &TemplateSt{}
	tt.LoadFile("./config/template/demo_001.yml")

	fmt.Printf("%+v", tt)
	link := "https://blog.sina.com.cn/s/blog_41772a550102zami.html?tj=1"

	//return
	doc := tt.Crawling(link)
	item, err := NewCompiler(doc).Parse(tt.Elements)
	fmt.Printf("%+v %+v", item, err)
}
