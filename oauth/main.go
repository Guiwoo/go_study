package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
)

func main() {

	clientId := os.Getenv("CLIENT_ID")
	secret := os.Getenv("SECRET")

	fmt.Println(clientId, secret)

	url := fmt.Sprintf("https://nid.naver.com/oauth2.0/authorize?"+
		"client_id=%s&"+
		"response_type=code&"+
		"redirect_uri=%s&"+
		"state=%s", clientId, "http://127.0.0.1:3000/callback", secret)

	fmt.Println(url)
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, fmt.Sprintf("<button><a href=\"%s\">네이버 로그인</a></button>", url))
	})

	type query struct {
		Code  string `query:"code"`
		State string `query:"state"`
	}

	e.GET("/callback", func(c echo.Context) error {
		var q query

		if err := c.Bind(&q); err != nil {
			log.Errorf("error is : ", err)
		}

		toeknUrl := fmt.Sprintf("https://nid.naver.com/oauth2.0/token?"+
			"client_id=%s&"+
			"client_secret=%s&"+
			"grant_type=authorization_code&"+
			"state=%s&"+
			"code=%s", clientId, secret, q.State, q.Code)

		resp, err := http.Get(toeknUrl)
		if err != nil {
			log.Errorf("Fail to get access token", err)
		}

		type tokenResponse struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			TokenTyep    string `json:"token_type"`
			ExpiredIn    int    `json:"expired_in"`
		}
		var token tokenResponse
		rs, _ := io.ReadAll(resp.Body)

		json.Unmarshal(rs, &token)
		return c.JSON(200, token)
	})

	if err := e.Start(":3000"); err != nil {
		panic(err)
	}
}
