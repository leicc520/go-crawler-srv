package parse

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
)

func TestError(t *testing.T) {
	e := ParseError{}
	err := errors.New("demo test")
	e.Wrapped("demo", err)
	e.Wrapped("demov2", err)

	fmt.Println(e, e.IsEmpty())
}

func TestReg(t *testing.T) {
	ss := "[\\d]+"
	result := "11adsf123dfdfdf"
	if reg, err := regexp.Compile(ss); err == nil {
		arrStr := reg.FindAllString(result, -1)
		if len(arrStr) > 0 {
			fmt.Println(arrStr)
		}
	}
}
