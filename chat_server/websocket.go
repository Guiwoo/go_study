package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type WebSocket struct {
	conn *websocket.Conn
}

func (w *WebSocket) setReadTimeout() error {
	return w.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
}

func newWebSocket(e echo.Context) (*WebSocket, error) {
	//room := e.Param("id")
	name := e.Request().Header.Get("name")

	var wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024 * 10,
		WriteBufferSize: 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := wsUpgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		return nil, err
	}

	w := &WebSocket{
		conn: conn,
	}

	w.conn.SetPongHandler(func(appData string) error {
		fmt.Println("pong handler called ")
		if err := w.setReadTimeout(); err != nil {
			w.conn.Close()
			return err
		}
		return nil
	})
	rm := GetRoomManager()
	user := NewUser(name, w)
	rm.EnterRoom(e.Param("id"), user)

	w.conn.SetCloseHandler(func(code int, text string) error {
		rm.SendExitSignal(w, e.Param("id"), name)
		return nil
	})
	return w, nil
}

func (w *WebSocket) writer() {

}

func (w *WebSocket) repeat() <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			var msg Message
			if err := w.conn.ReadJSON(&msg); err != nil {
				return
			}
		}
	}()
	return valueStream
}

func (w *WebSocket) receiver() {
	for {
		var msg Message
		if err := w.conn.ReadJSON(&msg); err != nil {
			return
		}

		rm := GetRoomManager()
		rm.SendChatSignal(w, msg.RoomNum, msg.UserId, msg.Message)
	}
}
