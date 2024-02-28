package api

import (
	"bytes"
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func Converter(html []byte, css []byte) {

	htmlPath := "/assets/tmp/index.html"
	cssPath := "/assets/tmp/index.css"
	if len(html) > 0 {
		createFile(htmlPath, html)
	}
	if len(css) > 0 {
		createFile(cssPath, css)
	}

	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte

	var quality = 200
	if err := chromedp.Run(ctx, fullScreenshot(`http://127.0.0.1:9000/assets/tmp/index.html`, quality, &buf)); err != nil {
		panic(err)
	}

	//buf 를 aws 로 쏘면 완벽
	UploadAWS_S3("guiwoo_test_css.png", bytes.NewReader(buf))
}

func createFile(path string, file []byte) {
	if err := os.WriteFile("/Users/guiwoopark/Desktop/personal/go_study/comfyui/"+path, file, 0o644); err != nil {
		log.Panicf("fail to create file %+v", err)
	}
}

func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
