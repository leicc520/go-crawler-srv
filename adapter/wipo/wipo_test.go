package wipo

import (
	"fmt"
	"testing"
)
import "github.com/Lazarus/lz-string-go"

func TestEncrypt(t *testing.T) {
	str := "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCAKoDKIANCAGYCWJAQgEoCCAcgCIgC+AujVo5I+Oo0ghyAFTbSK1EERJsAwtMUoA9iQDyLPv140kWgO4woANgAMNGAWQFrdughRwCl0SDY8jNAQAnmDEkgBGSAgAdgAmigA2CCRw0YpYEiAATAC0AKJgtABStADsACwAGhy6ALIJ5EEAzCzSHAAeZrQsNu157eFgAFpWBFnlugCuAJoAXuUA5nAsRXkAvOmTcFtQAIw0APoku7sAnE1ZfEAAA=="
	res, err := LZString.Decompress(str, "")
	fmt.Println(res, err)
	
	str = "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCLgEwAMZZAtBQJx30AqFFkbHFAWgATMB5XpWpMmzMgGZIAVnqz63ALogANCABmASxIAlACJqQAEwIkRtBkxABfJTZvqCATzDFIIAEZIEAO2NGADYIJHC+Rlg6HgCMNM4aAB4AHADuABIAVhoA5gCyMABe9ACungBSCAD2AJIpzPoAqsyBvmSVAGrG0QAyFGlJSAD0KQAKGkm+DQC8EcVwc1DR6gD6JNHR9JJktkAAA="
	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)
	
	str = "N4IgDiBcoM4KYEMBOBjAFlWBHKBtUALnFCLgEwAMZZAtBQJx30AqFFkbHFAWgATMB5XpWpMmzMgGZIAVnqz63ALogANCABmASxIAlACJqQAEwIkRtBkxABfJTfVIA9gHcYUAGwV1MAsjOQZAAsFA4gBACeYMSQIABGSAgAdsZGADYIJHBJRlg6sQCMNBEaAB4AHC4AEgBWGgDmALIwAF70AK5xAFIITgCSLsz6AKrMaUlkTgBqxgUAMhRV5UgA9C4AChrlScMAvLntcIdQBeoA+iQFBfSSZLZAA="
	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)
	
	str = "N4IgLgngDgpiBcIBGAnAhgOwCYgDQgBs0EQYM8QBHASxIAYBaARQFYAzANgFkAOADQBMAZQBeHABYsAbgGYAHgGMpQugAUA1nSwcAKgDEA5gC09RgOooAtknVcA7jICaAEQ50AvBUoBXGL4QAjPgA+iQBAQCcMgIgAL5AA=="
	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)
	
	str = "N4IgLgngDgpiBcIBGAnAhgOwCYgDQgBs0EQYM8QBHASxIAYBaAUQFUARagKQGkBhADwDOAKwBawqAEEASgEYAZgHtZAFQCa/AOoAZFAFYAxqLQAmAMwwoACQDUbGwBZFKgQA4AvBUoBXGL4Sy+AD6JLKyAJxmJiAAvkAAA=="
	res, err = LZString.Decompress(str, "")
	fmt.Println(res, err)
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