package api

import (
	"bytes"
	"comfyui/socket"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func UploadImage(url string) error {
	host := "http://127.0.0.1:8188/upload/image"

	buf, err := os.Open(url)
	if err != nil {
		return err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("image", "guiwoo.png")
	if err != nil {
		return err
	}

	_, err = io.Copy(fw, buf)
	if err != nil {
		log.Fatal(err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", host, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{
		// Set timeout to not be at mercy of microservice to respond and stall the server
		Timeout: time.Second * 20,
	}

	rsp, err := client.Do(req)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("fail to request %+v", rsp.StatusCode)
	}
	return nil
}

func CreateImage() {
	host := "http://127.0.0.1:8188/prompt"

	//하드코딩 변경 필요
	if err := UploadImage("/Users/guiwoopark/Desktop/personal/study/comfyui/assets/example2.png"); err != nil {
		panic(err)
	}

	fmt.Println("upload image done")

	// 하드코딩 변경 필요
	data, err := os.ReadFile("/Users/guiwoopark/Desktop/personal/study/comfyui/etc.json")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(data)

	req, _ := http.NewRequest("POST", host, buf)
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

func CreateImageWssConnect() <-chan string {
	c := make(chan string)
	CreateImage()
	go socket.Connect(c)
	return c
}
