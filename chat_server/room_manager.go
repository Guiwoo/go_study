package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var instance *RoomManager

type BroadCaster int

type RoomManager struct {
	Rooms       sync.Map
	broadcaster chan BroadCaster
}

func (rm *RoomManager) EnterRoom(user *User) {
	v, ok := rm.Rooms.LoadOrStore(user.roomNum, NewRoom(user.roomNum))
	if !ok {
		go v.(*Room).Run()
	}
	room := v.(*Room)
	fmt.Println("Enter someone", user)
	room.users = append(room.users, user)
	rm.SendEnterSignal(room, user.userId)
}

func (rm *RoomManager) SendEnterSignal(room *Room, userId string) {
	go func() {
		msg := NewMessage(room.number, userId, "님이 입장하셨습니다.")
		room.signal <- NewRoomSignal(msg, ENTER)
	}()
}

func (rm *RoomManager) SendExitSignal(exit *User) {
	go func() {
		defer func() {
			exit.ws.rw.RLock()
			err := exit.ws.conn.Close()
			if err != nil {
				fmt.Printf("error occur on send exit signal on socket close signal : %v\n", err)
			}
			exit.ws.isClose = true
			exit.ws.rw.RUnlock()
		}()
		msg := NewMessage(exit.roomNum, exit.userId, "님이 퇴장하셨습니다.")
		if value, ok := rm.Rooms.Load(exit.roomNum); ok {
			v := value.(*Room)

			v.mu.Lock()
			defer v.mu.Unlock()

			idx := -1
			for i, user := range v.users {
				if user.userId == exit.userId {
					idx = i
					break
				}
			}

			if idx == len(v.users) {
				v.users = v.users[:idx]
			} else if idx != -1 {
				tmp := v.users[:idx]
				v.users = append(tmp, v.users[idx+1:]...)
			}
			fmt.Println("current user list : ", v.users, v.number)
			v.signal <- NewRoomSignal(msg, EXIT)
		}
	}()
}

func (rm *RoomManager) SendChatSignal(w *WebSocket, param string, name string, message string) {
	go func() {
		msg := NewMessage(param, name, message)
		if value, ok := rm.Rooms.Load(param); ok {
			v := value.(*Room)
			v.signal <- NewRoomSignal(msg, CHAT)
		}
	}()
}
func (rm *RoomManager) BroadCaster() {
	one := time.NewTicker(45 * time.Second)
	two := time.NewTicker(90 * time.Second)
	for {
		select {
		case <-one.C:
			rm.BroadcastingUser()
		case <-two.C:
			rm.BroadcastingChatHistory()
		}
	}
}
func (rm *RoomManager) BroadcastingUser() {
	rm.Rooms.Range(func(key, value interface{}) bool {
		room := value.(*Room)
		fmt.Println(len(room.users), room.number)
		users := make([]string, len(room.users))
		for i, v := range room.users {
			users[i] = v.userId
		}
		msg := NewMessage(room.number, "🍕 System", "현재 접속자 : "+strings.Join(users, " ")+"🍕")
		room.signal <- NewRoomSignal(msg, BroadCast)
		return true
	})
}

func (rm *RoomManager) BroadcastingChatHistory() {
	rm.Rooms.Range(func(key, value interface{}) bool {
		room := value.(*Room)
		link := room.link
		history := link.GenerateHistory()
		msg := NewMessage(room.number, "✅ System", history)
		room.signal <- NewRoomSignal(msg, BroadCast)
		return true
	})
}

func GetRoomManager() *RoomManager {
	return instance
}

func NewRoomManager() *RoomManager {
	c := make(chan BroadCaster)
	instance = &RoomManager{broadcaster: c}
	go instance.BroadCaster()
	return instance
}
