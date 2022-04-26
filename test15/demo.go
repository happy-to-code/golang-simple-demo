// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	// var example string
	for i := 0; i < 10; i++ {

		err := chromedp.Run(ctx,
			chromedp.Navigate(`https://www.jszzb.gov.cn/col22/81608.html`),
			// wait for footer element is visible (ie, page is loaded)
			chromedp.WaitVisible(`#196a21d2-42cb-48f8-b841-c638da7a9694`, chromedp.ByID),
			// find and click "Expand All" link
			// chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
			// // retrieve the value of the textarea
			// chromedp.Value(`#example_After .play .input textarea`, &example),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	// log.Printf("Go's time.After example:\n%s", example)
}
