package socket

import (
	"comfyui/types"
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

func readClientMsg(ws *websocket.Conn, done chan bool) <-chan types.ComfySocketResp {
	stream := make(chan types.ComfySocketResp)
	go func() {
		defer close(stream)
		defer fmt.Println("done go routine")
		for {
			var message types.ComfySocketResp
			err := websocket.JSON.Receive(ws, &message)
			if err != nil {
				fmt.Printf("fail to get messgae from ws %+v\n", err)
			}
			select {
			case <-done:
				return
			case stream <- message:
			}
		}
	}()
	return stream
}

func Connect(sign chan bool) {
	ws, err := websocket.Dial(
		"ws://127.0.0.1:8188/ws?clientId=87655da4394a4f1f96301f9f8933b114",
		"",
		"http://127.0.0.1:8188")
	if err != nil {
		fmt.Printf("fail to connect ws connection %+v\n", err)
		return
	}

	done := make(chan bool)
	msg := readClientMsg(ws, done)
	defer func() {
		close(done)
		ws.Close()
		sign <- true
	}()
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("pass the time")
		case msg, _ := <-msg:
			fmt.Printf("msg %+v\n", msg)
			if msg.Type == "executed" {
				done <- true
				fmt.Println("[info] socket connections break created Images")
				return
			}
		}
	}
}