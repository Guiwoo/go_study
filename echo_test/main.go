package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func callEcho() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		fmt.Println("sibal")
		return c.JSON(200, struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}{"testing", "From Server !"})
	})

	//s := http.Server{
	//	Addr:    ":4000",
	//	Handler: e,
	//	TLSConfig: &tls.Config{
	//		NextProtos: []string{acme.ALPNProto},
	//	},
	//}
	//log.Error(s.ListenAndServeTLS("/Users/guiwoopark/Documents/cert.pem", "/Users/guiwoopark/Documents/key.pem"))
	log.Error(e.Start(":4006"))
}

func main() {
	callEcho()
}
