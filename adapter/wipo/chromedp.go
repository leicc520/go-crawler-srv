package wipo

import (
	"context"
	"github.com/chromedp/chromedp"
	"math/rand"
	"time"
)

func closeByDialog(dialog string, cancel context.CancelFunc) {
	if dialog == "" {
		return
	}
	cancel()
}

func setCountry(country string, ctx context.Context) {
	chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Click("#ui-id-6"),
		chromedp.Sleep(time.Duration(rand.Intn(11) + rand.Intn(11)*1000*1000*1000)),
		chromedp.SendKeys("#OO_input", country),
		chromedp.Click("#country_search > div.searchButtonContainer.bottom.right > a"),
	})
}

func setDate(applicationDate string, ctx context.Context) {
	chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Click("#ui-id-4"),
		chromedp.Sleep(time.Duration(rand.Intn(11) + rand.Intn(11)*1000*1000*1000)),
		chromedp.SendKeys("#AD_input", applicationDate),
		chromedp.Click("#country_search > div.searchButtonContainer.bottom.right > a"),
	})
}

