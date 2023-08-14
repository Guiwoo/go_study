package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func main() {

	//로그가 찍힐때 마다 계속 조회 해서 시간 값 비교해서 넘기고
	// 그다음으로 넘기고

	logger := log.New(os.Stdout, "[something.customLog]", log.LstdFlags|log.Ltime)
	logger.Println("Start echo")
	logger.SetOutput()

	e := echo.New()

	e.GET("/test", func(c echo.Context) error {
		logger.Printf("Got [Request : %v][Ip : %s] [Headers : %v] [URL : %v] [Body : %v]",
			c.Request(), c.Request().RemoteAddr, c.Request().Header, c.Request().URL, c.Request().GetBody)
		return c.JSON(200, "work as perfect")
	})

	server := http.Server{
		Addr:     ":8080",
		Handler:  e,
		ErrorLog: logger,
	}

	logger.Printf("error comes here %+v", server.ListenAndServe())
}
