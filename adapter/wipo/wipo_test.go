package wipo

import (
	"encoding/json"
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
	"testing"
	"time"
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
	
	str = "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCLgEwAMZZAtBQJx30AqFFkbHFAWgATMB5XpWpMmzMgGZIAVnqz63ALogANCABmASxIAlACJqQAEwIkRtBkxABfJTfVIA9gHcYUAGwV1MAsjOQAIwAHBQOIAQAnmDEkCAARkgIAHbGRgA2CCRwyUZYOnHBNAAKAOaSHggoHgCaTnAUAMIaCABugUhoMkgAsnAAVgBq9ACqNfQwAMouAOyRga4AcggAYgBeTszBALx5AK5wB1CB6gD6JIGB9JJktkAAA="
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
	ss := &WipoSt{}
	ss.Run("2022-09-09", "2022-09-12")
}
