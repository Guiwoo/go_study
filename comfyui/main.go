package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type Prompt struct {
	Msg       string `json:"prompt"`
	Translate string `json:"translate"`
}

func callNaverAPI() (string, error) {
	host := "https://openapi.naver.com/v1/papago/n2mt"

	data := url.Values{}
	data.Set("source", "en")
	data.Set("target", "ko")
	data.Set("text", "Nice to meet you. BTS")

	req, err := http.NewRequest("POST", host, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Naver-Client-Id", "Nr6PZ2Nj7H2n2pKur5vJ")
	req.Header.Set("X-Naver-Client-Secret", "8kny84_ER9")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	fmt.Println(string(body))

	return "", nil
}

func translate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 400)
		return
	}
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "wrong data", 400)
		return
	}
	var body Prompt
	if err := json.Unmarshal(resp, &body); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to unmarshal data", 400)
		return
	}
	body.Translate = "호잇"
	jsonResp, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "fail to marshaling", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}

func Render() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fp := path.Join("/Users/guiwoopark/Desktop/personal/study/comfyui/templates", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			log.Fatalf("fail to open file %+v", err)
		}
		if err = tmpl.Execute(w, nil); err != nil {
			log.Fatalf("fail to exute template %+v", err)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/translate", translate)
	port := ":9000"
	http.Handle("/js", http.FileServer(http.Dir("/Users/guiwoopark/Desktop/personal/study/comfyui/templates/js")))
	http.ListenAndServe(port, mux)
}

func createImage() {
	host := "http://127.0.0.1:8188/prompt"
	done := make(chan interface{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("finished generated image")
				break
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("Creating...")
			}
		}
	}()

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

	done <- "got response"

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	Render()
}
