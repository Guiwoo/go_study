package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"runtime"
	"sync"
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

	var wsUpgrade = websocket.Upgrader{
		ReadBufferSize:  1024 * 10,
		WriteBufferSize: 1024 * 10,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := wsUpgrade.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		return nil, err
	}

	w := &WebSocket{
		conn: conn,
	}

	w.conn.SetPongHandler(func(appData string) error {
		fmt.Println("pong handler called ")
		if err := w.setReadTimeout(); err != nil {
			log.Errorf("error is on pong Handler", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return nil
			} else {
				w.conn.Close()
			}
			return nil
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

func (w *WebSocket) repeat(done chan interface{}) <-chan *Message {
	valueStream := make(chan *Message)
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			default:
				var msg Message
				if err := w.conn.ReadJSON(&msg); err != nil {
					log.Errorf("error is : ", err)
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						return
					}
					continue
				}
				valueStream <- &msg
			}
		}
	}()
	return valueStream
}

func (w *WebSocket) fanout(done chan interface{}) []<-chan *Message {
	cpus := runtime.NumCPU()
	chans := make([]<-chan *Message, cpus)

	for i := 0; i < cpus; i++ {
		chans[i] = w.repeat(done)
	}
	return chans
}

func orDone(done <-chan interface{}, msg <-chan *Message) <-chan *Message {
	valueStream := make(chan *Message)
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case x, ok := <-msg:
				if !ok {
					return
				}
				select {
				case valueStream <- x:
				case <-done:
				}
			}
		}
	}()
	return valueStream
}

func fanIn(done <-chan interface{}, channels ...<-chan *Message) <-chan *Message {
	var wg sync.WaitGroup
	multiplexedStream := make(chan *Message)

	multiple := func(c <-chan *Message) {
		defer wg.Done()
		for a := range orDone(done, c) {
			select {
			case <-done:
			case multiplexedStream <- a:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiple(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func (w *WebSocket) receiver() {
	done := make(chan interface{})
	defer close(done)
	messageStream := fanIn(done, w.fanout(done)...)

	for msg := range messageStream {
		rm := GetRoomManager()
		rm.SendChatSignal(w, msg.RoomNum, msg.UserId, msg.Message)
	}
}
