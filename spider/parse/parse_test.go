package parse

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	e := ParseError{}
	err := errors.New("demo test")
	e.Wrapped("demo", err)
	e.Wrapped("demov2", err)

	fmt.Println(e, e.IsEmpty())
}
