package main

import (
	"bytes"
	"comfyui/api"
	"comfyui/types"
	"encoding/json"
	"fmt"
	sc "github.com/gorilla/websocket"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"time"
)

var upgrader = sc.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Translate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 400)
		return
	}
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "wrong data", 400)
		return
	}
	var body types.Prompt
	if err := json.Unmarshal(resp, &body); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to unmarshal data", 400)
		return
	}

	pTranslated, err := api.CallPapagoAPI(body.Positive)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "fail to call naver api", 400)
		return
	}

	nTranslated, err := api.CallPapagoAPI(body.Negative)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "fail to call naver api", 400)
		return
	}

	body.PTranslate = pTranslated
	body.NTranslate = nTranslated
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

func CallDalle(w http.ResponseWriter, r *http.Request) {
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "wrong data", 400)
		return
	}
	var body types.Prompt
	if err := json.Unmarshal(resp, &body); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to unmarshal data", 400)
		return
	}
	path, err := api.CallDalle(body.Positive, body.ClientID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "fail to call dalle", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{
		"msg": "ok",
		"url": path,
	}
	value, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "fail to marshal", 400)
		return
	}
	w.Write(value)
	return
}

func QueuePrompt(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	var msg types.QueueRequest
	if err := c.ReadJSON(&msg); err != nil {
		fmt.Println("json parse error", err)
	}

	getImage := api.CreateImageWssConnect(msg)

	for {
		select {
		case <-time.After(1 * time.Second):
			c.WriteMessage(sc.TextMessage, []byte("progress"))
		case filename := <-getImage:
			bucketName := "test-guiwoo"
			awsFullPath := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)
			c.WriteMessage(sc.TextMessage, []byte(awsFullPath))
			return
		}
	}
}

func Gpt(w http.ResponseWriter, r *http.Request) {
	resp, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "wrong data", 400)
		return
	}
	var body types.GPTRequest
	if err := json.Unmarshal(resp, &body); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to unmarshal json", 400)
		return
	}

	data, err := api.CallGPT(body.Input)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "fail to call api gpt", 400)
		return
	}

	rsp, _ := json.Marshal(&data)

	w.WriteHeader(200)
	w.Write(rsp)
}

func HtmlConverter(w http.ResponseWriter, r *http.Request) {
	htmlFile, _, _ := r.FormFile("html")
	cssFile, _, _ := r.FormFile("css")
	var (
		htmlData = bytes.NewBuffer(make([]byte, 0))
		cssData  = bytes.NewBuffer(make([]byte, 0))
	)
	if htmlFile != nil {
		io.Copy(htmlData, htmlFile)
	}
	if cssFile != nil {
		io.Copy(cssData, cssFile)
	}

	api.Converter(htmlData.Bytes(), cssData.Bytes())
}

func Render() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fp := path.Join("/Users/guiwoopark/Desktop/personal/go_study/comfyui/templates", "index.html")
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
	mux.HandleFunc("/gpt", Gpt)
	mux.HandleFunc("/translate", Translate)
	mux.HandleFunc("/dalle", CallDalle)
	mux.HandleFunc("/generate", QueuePrompt)
	mux.HandleFunc("/converter", HtmlConverter)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/Users/guiwoopark/Desktop/personal/go_study/comfyui/assets"))))
	mux.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("/Users/guiwoopark/Desktop/personal/go_study/comfyui/output"))))
	port := ":9000"
	http.ListenAndServe(port, mux)
}

func main() {
	Render()
}
