package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type SearchKeyword struct {
	Title      string `json:"title"`
	IsUse      bool   `json:"is_use"`
	IsPrivate  bool   `json:"is_private"`
	ContentsId int    `json:"contents_id"`
	Order      string `json:"order"`
}

func callEcho() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		var keyword SearchKeyword
		if err := c.Bind(&keyword); err != nil {
			fmt.Printf("got error ", err)
		}
		fmt.Println(keyword)
		fmt.Println("got request")
		return c.JSON(200, struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}{"testing", "From Server ! Here is different server on 9096"})
	})

	data, err := http.Get("https://localhost:9098")
	if err != nil {
		log.Error(err)
	} else {
		log.Infof("got data %v\n", data)
	}
	log.Error(e.Start(":9096"))
}

func main() {
	callEcho()
}
