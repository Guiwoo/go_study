package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

var (
	cert string = "/Users/guiwoopark/Documents/cert.pem"
	key  string = "/Users/guiwoopark/Documents/key.pem"

	serverCert = "/home/kbp/etc/ssl/cert.pem"
	serverKey  = "/home/kbp/etc/ssl/key.pem"

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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	//s := http.Server{
	//	Addr:    ":9098",
	//	Handler: e,
	//	TLSConfig: &tls.Config{
	//		NextProtos: []string{acme.ALPNProto},
	//	},
	//}
	//customLog.Error(s.ListenAndServeTLS(serverCert, serverKey))

	http.Get("https://im.plea.kr:9098/test")
	log.Error(e.Start(":9098"))
}
