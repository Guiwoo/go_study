package main

import (
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
	if err := api.CallDalle(body.Positive); err != nil {
		fmt.Println(err)
		http.Error(w, "fail to call dalle", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := map[string]interface{}{
		"msg": "ok",
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
			c.WriteMessage(sc.TextMessage, []byte(filename))
			return
		}
	}
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
	mux.HandleFunc("/translate", Translate)
	mux.HandleFunc("/dalle", CallDalle)
	mux.HandleFunc("/generate", QueuePrompt)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/Users/guiwoopark/Desktop/personal/study/comfyui/assets"))))
	mux.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("/Users/guiwoopark/Desktop/personal/study/comfyui/output"))))
	port := ":9000"
	http.ListenAndServe(port, mux)
}

func main() {
	Render()
}
