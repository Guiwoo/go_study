package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	if err := chromedp.Run(ctx, elementScreenshot(`https://www.lge.co.kr/lg-styler/sc5mbr60`, `div#overview.container-fluid.iw_section`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("LG가전_스타일오브제.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	// capture entire browser viewport, returning png with quality=90
	//var quality = 200
	//if err := chromedp.Run(ctx, fullScreenshot(`http://127.0.0.1:9000`, quality, &buf)); err != nil {
	//	log.Fatal(err)
	//}
	//if err := os.WriteFile(fmt.Sprintf("fullScreenshot_%d.png", quality), buf, 0o644); err != nil {
	//	log.Fatal(err)
	//}

	log.Printf("wrote LG가전_스타일오브제.png and fullScreenshot_200.png")
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
