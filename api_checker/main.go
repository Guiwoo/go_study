package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
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
	Data   string `json:"data,omitempty"`
	Url    string `json:"url,omitempty"`
}

func getEnv(env string) string {
	value, ok := os.LookupEnv(env)
	if ok == false {
		log.Fatalf("fail to get env %+v", env)
		return ""
	}
	return value
}

func getImgFileToBase64(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(file)
}

func ocrSendURL() {
	host := getEnv("naver_url")
	secret := getEnv("xor_secret")

	data := map[string]interface{}{
		"version":   "V2",
		"requestId": uuid.NewString(),
		"timestamp": time.Now().Nanosecond(),
		"lang":      "ko",
		"images": []DATA{
			DATA{
				Format: "png",
				Name:   "test-guiwoo",
				Url:    "https://user-images.githubusercontent.com/67041069/201728725-6611c514-e1a5-4d78-9060-8965be25fd1c.png",
			},
		},
		"resultType": "string",
	}

	js, _ := json.Marshal(&data)

	req, err := http.NewRequest("POST", host, bytes.NewReader(js))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-OCR-SECRET", secret)

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

func ocrSendData() {
	host := getEnv("naver_url")
	secret := getEnv("xor_secret")

	data := map[string]interface{}{
		"version":   "V2",
		"requestId": uuid.NewString(),
		"timestamp": time.Now().Nanosecond(),
		"lang":      "ko",
		"images": []DATA{
			DATA{
				Format: "png",
				Name:   "test-guiwoo",
				Data:   getImgFileToBase64("/Users/guiwoopark/Desktop/something.png"),
			},
		},
		"resultType": "string",
	}

	js, _ := json.Marshal(&data)

	req, err := http.NewRequest("POST", host, bytes.NewReader(js))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-OCR-SECRET", secret)

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

type MultiPartFormData struct {
	Version   string      `json:"version"`
	RequestID string      `json:"requestId"`
	TimeStamp int         `json:"timestamp"`
	Images    []ImageData `json:"images"`
}
type ImageData struct {
	Format string `json:"format"`
	Name   string `json:"name"`
}

func ocrSendFile() {
	host := getEnv("naver_url")
	secret := getEnv("xor_secret")
	fmt.Println(host, secret)

	file, err := os.Open("/Users/guiwoopark/Desktop/something.png")
	defer file.Close()
	if err != nil {
		log.Panicf("fail to open file %+v", err)
	}
	data := MultiPartFormData{
		Version:   "V2",
		RequestID: uuid.NewString(),
		TimeStamp: time.Now().Nanosecond(),
		Images: []ImageData{
			{
				Format: "png",
				Name:   "test-guiwoo",
			},
		},
	}
	value, _ := json.Marshal(&data)

	reqBody := bytes.Buffer{}
	writer := multipart.NewWriter(&reqBody)
	if err := writer.WriteField("message", string(value)); err != nil {
		panic(err)
	}
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if _, err := io.Copy(part, file); err != nil {
		panic(err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", host, &reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-OCR-SECRET", secret)

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

func main() {
	ocrSendFile()
}
