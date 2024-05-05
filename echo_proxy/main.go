package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
)

var (
	cert string = "/Users/guiwoopark/Documents/cert.pem"
	key  string = "/Users/guiwoopark/Documents/key.pem"

	serverCert = "/home/kbp/dynamo_util/ssl/cert.pem"
	serverKey  = "/home/kbp/dynamo_util/ssl/key.pem"

	targetServer  = "http://im.plea.kr:9096"
	currentServer = "/test"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(200, struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}{"testing", "From Server !"})
	})

	url1, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{
			URL: url1,
		},
	})))

	go func() {
		e2 := echo.New()
		e2.GET("/test", func(c echo.Context) error {
			return c.JSON(200, "this is test")
		})
		e.Start(":8081")
	}()

	log.Error(e.Start(":9098"))
}
