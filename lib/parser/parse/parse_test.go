package parse

import (
	"errors"
	"fmt"
	"github.com/leicc520/go-crawler-srv/lib"
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

func TestQuery(t *testing.T) {
	str := `<div class="a-box-inner a-padding-medium"><!-- Detailed Seller Information -->
                <div class="a-row a-spacing-small"><h3>Detailed Seller Information</h3></div><div class="a-row a-spacing-none"><span class="a-text-bold">Business Name:
                            </span><span>Corso Bale.Inc</span></div><div class="a-row a-spacing-none"><span class="a-text-bold">Business Address:
                            </span></div><div class="a-row a-spacing-none indent-left"><span>5822 W Third ST #101</span></div><div class="a-row a-spacing-none indent-left"><span>Los Angeles</span></div><div class="a-row a-spacing-none indent-left"><span>CA</span></div><div class="a-row a-spacing-none indent-left"><span>90036</span></div><div class="a-row a-spacing-none indent-left"><span>US</span></div><!-- Detailed Seller Information -->
            </div>`
	tt, err := NewQueryParse(str)
	fmt.Println(tt, err)
	astr, err := tt.InnerTexts(".indent-left span")
	fmt.Println(astr, err)

	str = lib.StripTags(str)
	fmt.Println(str)
}

func TestTable(t *testing.T) {
	str := `<table>
<tr><th></th><th class="a-text-right">30 days</th>
<th class="a-text-right">90 days</th>
<th class="a-text-right">12 months</th>
<th class="a-text-right">Lifetime</th></tr>

<tr><td class="a-nowrap" style="width:1px;">Positive</td> <td class="a-text-right"> <span class="a-color-success">63</span>% </td> <td class="a-text-right"> 
<span class="a-color-success">68</span>% </td> <td class="a-text-right"> <span class="a-color-success">75</span>% </td> <td class="a-text-right"> 
<span class="a-color-success">89</span>% </td></tr>
<tr><td class="a-nowrap" style="width:1px;">Neutral</td> <td class="a-text-right"> <span class="a-color-secondary">3</span>% </td> <td class="a-text-right"> 
<span class="a-color-secondary">1</span>% </td> <td class="a-text-right"> <span class="a-color-secondary">3</span>% </td> <td class="a-text-right">
<span class="a-color-secondary">2</span>% </td></tr><tr><td class="a-nowrap" style="width:1px;">Negative</td> <td class="a-text-right">
<span class="a-color-error">34</span>% </td> <td class="a-text-right"> <span class="a-color-error">31</span>% </td> <td class="a-text-right"> 
<span class="a-color-error">22</span>% </td> <td class="a-text-right"> <span class="a-color-error">9</span>% </td></tr>
<tr><td class="a-nowrap" style="width:1px;">Count</td><td class="a-text-right"><span>38</span></td><td class="a-text-right"><span>120</span>
</td><td class="a-text-right"><span>549</span></td><td class="a-text-right"><span>12,820</span></td>
</tr></table>`
	fmt.Println(str)

	tt, err := NewQueryParse(str)
	fmt.Println(tt, err)
	astr, err := tt.InnerTexts("table tr")
	fmt.Println(astr, err)

	expr := `//table[@id='feedback-summary-table']`
	ok, err := regexp.MatchString(`//table\[[^\]]+\]`, expr)
	fmt.Println(ok, err)
}
