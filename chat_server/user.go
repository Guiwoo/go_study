package main

import "fmt"

type User struct {
	userId  string
	roomNum string
	ws      *WebSocket
}

func NewUser(userId, roomNum string, ws *WebSocket) *User {
	fmt.Println("Socket is created", ws)
	return &User{userId, roomNum, ws}
}
