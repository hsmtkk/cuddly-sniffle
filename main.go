package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

const CME_URL = "https://www.moomoo.com/ja/stock/XW2407-US?"

func main() {
	chromeContext, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string

	err := chromedp.Run(chromeContext,
		chromedp.Navigate(CME_URL),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			res, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}))
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("page.html", []byte(res), 0644); err != nil {
		log.Fatal(err)
	}
}
