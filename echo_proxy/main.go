package main

import (
	"crypto/tls"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme"
	"net/http"
	"net/url"
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

	url1, _ := url.Parse("http://localhost:9096")
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	grp := e.Group(currentServer)
	grp.Use(middleware.Recover())
	grp.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer:  middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{URL: url1}}),
		Transport: transport,
	}))

	s := http.Server{
		Addr:    ":9098",
		Handler: e,
		TLSConfig: &tls.Config{
			NextProtos: []string{acme.ALPNProto},
		},
	}
	log.Error(s.ListenAndServeTLS(cert, key))
}
