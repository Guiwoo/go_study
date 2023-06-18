package main

import (
	"github.com/labstack/echo/v4"
)

func startChat(e echo.Context) error {
	socket, err := newWebSocket(e)
	if err != nil {
		return err
	}
	go socket.writer()
	go socket.receiver()

	return nil
}

// 웻소켓 세션 담당
func main() {
	e := echo.New()
	NewRoomManager()

	e.GET("/chat/live/:id", startChat)
	e.Logger.Fatal(e.Start(":1234"))
}
