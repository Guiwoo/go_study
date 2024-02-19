package api

import (
	"bytes"
	"comfyui/socket"
	"comfyui/types"
	"encoding/json"
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
	//todo cliendId + .png 로 변경
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

func configJson(data []byte, req types.QueueRequest) ([]byte, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	//ClientID
	SetClientID(req.ClientID, m)
	//Model
	SetStringJson(model, modelKey, req.Model, m)
	//Positive
	SetStringJson(positive, positiveKey, req.Positive, m)
	//Negative
	SetStringJson(negative, negativeKey, req.Negative, m)
	//Seed
	SetBigInt(ksampler, ksamplerSeed, req.Seed, m)
	//CFG
	SetIntJson(ksampler, ksamplerCfg, req.Cfg, m)
	//Setps
	SetIntJson(ksampler, ksamplerSteps, req.Steps, m)
	//Width
	SetIntJson(image, imageWidth, req.Width, m)
	//Height
	SetIntJson(image, imageHeight, req.Height, m)
	//BatchSize
	SetIntJson(image, imageBatchSize, req.BatchSize, m)
	//SetOutput
	SetStringJson(output, fileNamePrefix, req.ClientID, m)

	jsonData, err := json.Marshal(&m)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func CreateImage(qReq types.QueueRequest) error {
	host := "http://127.0.0.1:8188/prompt"

	// todo client id로 변경하기 imageURL 확인해서 보내기 , s3 연동해보기
	if err := UploadImage("/Users/guiwoopark/Desktop/personal/study/comfyui/assets/example2.png"); err != nil {
		panic(err)
	}

	fmt.Println("upload image done")

	data, err := os.ReadFile("/Users/guiwoopark/Desktop/personal/study/comfyui/etc.json")
	if err != nil {
		panic(err)
	}
	jsonData, err := configJson(data, qReq)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(jsonData)

	req, _ := http.NewRequest("POST", host, buf)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}
	//todo history 호출해보기 완료된건지 확인해야하는데... 응답값 돌려주고 해당 응답값들고 넘겨버려 c에다가 데이터 넣어주기
	fmt.Println(string(body))
	//if strings.Contains(string(body), "error") {
	//	return fmt.Errorf("fail to request %+v", string(body))
	//}
	return nil
}

func CreateImageWssConnect(data types.QueueRequest) <-chan string {
	c := make(chan string)
	if err := CreateImage(data); err != nil {
		log.Fatalf("err %+v\n", err)
	}

	go socket.Connect(c, data.ClientID)

	return c
}
