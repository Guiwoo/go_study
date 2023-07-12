package main

import (
	"crypto/tls"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(200, struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}{"testing", "From Server !"})
	})

	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().TLS != nil {
				redirectURL := "http://" + "localhost:9086" + c.Request().URL.String()
				return c.Redirect(http.StatusMovedPermanently, redirectURL)
			}
			return next(c)
		}
	})

	go func() {
		s := http.Server{
			Addr:    ":9087",
			Handler: e,
			TLSConfig: &tls.Config{
				NextProtos: []string{acme.ALPNProto},
			},
		}
		log.Error(s.ListenAndServeTLS("/Users/guiwoopark/Documents/cert.pem", "/Users/guiwoopark/Documents/key.pem"))
	}()

	log.Error(e.Start(":9086"))
}
