package dbsrv

import (
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
	"github.com/leicc520/go-crawler-srv/service/dbsrv/spider/models"
	"github.com/leicc520/go-orm"
	_ "github.com/lib/pq"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func init() {
	os.Chdir("../../")
	lib.InitConfig()
}

func TestOrmDemo(t *testing.T) {
	sorm   := models.NewSpAmazonAsin()
	aotiid := sorm.Save(1, orm.SqlMap{"updated_at":time.Now()})
	fmt.Println(aotiid)
	dd := models.SpAmazonAsinSt{}
	data1 := sorm.NoCache().GetOne(1).ToStruct(&dd)

	fmt.Println(data1, dd)

	fields := orm.SqlMap{"asin":"11111", "updated_at":time.Now(), "created_at":time.Now()}
	autoid := sorm.NewOne(fields, nil)
	fmt.Println(autoid)

	data := models.SpAmazonAsinSt{}
	err := sorm.NoCache().GetOne(autoid).ToStruct(&data)
	fmt.Println(err, data)
	re1 := sorm.Save(autoid, orm.SqlMap{"asin":"22222"})
	col := sorm.GetColumn(0, -1, nil, "asin", "id", orm.ASC)
	ls1 := sorm.GetList(0, -1, func(st *orm.QuerySt) string {
		st.Where("id", autoid)
		st.OrderBy("id", orm.DESC)
		return st.GetWheres()
	}, "id,asin,created_at")
	mp1 := sorm.GetAsMap(0, -1, nil, "id,asin")
	mp2 := sorm.GetNameMap(0, -1, nil, "id,asin,updated_at", "id")
	re2 := sorm.Delete(autoid)
	fmt.Println(re1, re2, col, ls1, mp1, mp2)
}

func TestName(t *testing.T) {
	gdir, _ := os.Getwd()
	gdir  = filepath.Join(gdir, "service", "dbsrv", "spider", "models")
	orm.CreatePGSQLModels("spiderdbmaster", "spiderdbslaver", gdir)
}
