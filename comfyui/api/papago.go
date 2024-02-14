package api

import (
	"comfyui/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func CallPapagoAPI(kor string) (string, error) {
	host := "https://openapi.naver.com/v1/papago/n2mt"
	id := getEnv("PapagoID")
	secret := getEnv("PapagoSecret")

	data := url.Values{}
	data.Set("source", "ko")
	data.Set("target", "en")
	data.Set("text", strings.TrimSpace(kor))

	req, err := http.NewRequest("POST", host, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Naver-Client-Id", id)
	req.Header.Set("X-Naver-Client-Secret", secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()
	var papago types.PapagoResp
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	if err := json.Unmarshal(body, &papago); err != nil {
		fmt.Println("error fail to unmarshal : ", err)
		return "", err
	}
	log.Printf("[info] got : %+v\n", string(body))

	return papago.Message.Result.TranslatedText, nil
}
