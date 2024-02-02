package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func papagoTest() {
	host := "https://openapi.naver.com/v1/papago/n2mt"

	data := url.Values{}
	data.Set("source", "en")
	data.Set("target", "ko")
	data.Set("text", "Nice to meet you. BTS")

	req, err := http.NewRequest("POST", host, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Naver-Client-Id", "Nr6PZ2Nj7H2n2pKur5vJ")
	req.Header.Set("X-Naver-Client-Secret", "8kny84_ER9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))
}

type DATA struct {
	Format string `json:"format"`
	Name   string `json:"name"`
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func main() {
	host := ""
	file, _ := os.ReadFile("/Users/guiwoopark/Desktop/something.png")

	encodedString := base64.StdEncoding.EncodeToString(file)

	data := map[string]interface{}{
		"version":   "V2",
		"requestId": uuid.NewString(),
		"timestamp": time.Now().Nanosecond(),
		"lang":      "ko",
		"images": []DATA{
			DATA{"png", "something"},
		},
		"enableTableDetection": true,
	}
	payload := map[string]interface{}{
		"message": data,
	}
	js, _ := json.Marshal(&payload)

	//req, err := http.NewRequest("POST", host, bytes.NewReader(js))
	req, err := newfileUploadRequest(host)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-OCR-SECRET", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))
}
