package main

import (
	"crypto/tls"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	cert string = "/Users/guiwoopark/Documents/cert.pem"
	key  string = "/Users/guiwoopark/Documents/key.pem"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(200, struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}{"testing", "From Server !"})
	})
	url1, _ := url.Parse("http://localhost:4006/test")
	proxy := httputil.NewSingleHostReverseProxy(url1)

	e.Any("/", func(c echo.Context) error {
		req := c.Request()

		// Modify request for proxy
		req.URL.Scheme = url1.Scheme
		req.URL.Host = url1.Host
		req.URL.Path = url1.Path
		req.Host = url1.Host
		proxy.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	s := http.Server{
		Addr:    ":4000",
		Handler: e,
		TLSConfig: &tls.Config{
			NextProtos: []string{acme.ALPNProto},
		},
	}
	log.Error(s.ListenAndServeTLS(cert, key))
}
