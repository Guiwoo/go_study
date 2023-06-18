package main

import "log"

type Room struct {
	users  []*User
	link   *MessageLink
	number string
	signal chan *RoomSignal
}

type Signal int

const (
	ENTER Signal = iota
	EXIT
	CHAT
	BroadCast
)

type RoomSignal struct {
	*Message
	Signal
}

func NewRoomSignal(msg *Message, signal Signal) *RoomSignal {
	return &RoomSignal{msg, signal}
}

func (r *Room) Run() {
	for {
		select {
		case msg := <-r.signal:
			switch msg.Signal {
			case ENTER:
				r.sendMessage(msg.Message)
			case EXIT:
				for i, v := range r.users {
					if v.userId == msg.UserId {
						r.users = append(r.users[:i], r.users[i+1:]...)
					}
				}
				r.sendMessage(msg.Message)
			case CHAT:
				r.link.Add(&MessageLink{msg: msg.Message})
				r.sendMessage(msg.Message)
			case BroadCast:
				r.sendMessage(msg.Message)
			}
		}
	}
}

func (r *Room) sendMessage(msg *Message) {
	go func() {
		for _, v := range r.users {
			if err := v.ws.conn.WriteJSON(msg.GenerateMessage()); err != nil {
				log.Println(err)
				break
			}
		}
	}()
}

func NewRoom(num string) *Room {
	c := make(chan *RoomSignal, 1000)
	return &Room{number: num, signal: c, link: NewMessageLink()}
}
