package main

import (
	"fmt"
	"strings"
)

type Message struct {
	RoomNum string `json:"roomNum"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
}

func (m *Message) GenerateMessage() string {
	return fmt.Sprintf("[%s][%s] : %s .", m.RoomNum, m.UserId, m.Message)
}

func NewMessage(roomNum, userId, message string) *Message {
	return &Message{
		RoomNum: roomNum,
		UserId:  userId,
		Message: message,
	}
}

type MessageLink struct {
	next *MessageLink
	msg  *Message
}

func (m *MessageLink) Add(link *MessageLink) {
	if m.next == nil {
		m.next = link
		return
	}
	m.next.Add(link)
}

func (m *MessageLink) GenerateHistory() string {
	sb := strings.Builder{}
	for m.next != nil {
		sb.WriteString(m.next.msg.GenerateMessage())
		m = m.next
	}
	return sb.String() + "✅"
}

func NewMessageLink() *MessageLink {
	return &MessageLink{}
}
