package main

type User struct {
	userId string
	ws     *WebSocket
}

func NewUser(userId string, ws *WebSocket) *User {
	return &User{userId, ws}
}
