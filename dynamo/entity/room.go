package entity

import (
	"context"
	"dynamo/dynamo_util"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"time"
)

const (
	RPREFIX = "Room#"
)

type Room struct {
	PK     string
	SK     string
	Entity dynamo_util.EntityType
	Title  string
	Date   time.Time
}

func roomCreate(title string) Room {
	core := dynamo_util.IdGenerator()
	pk := RPREFIX + core
	sk := RPREFIX + Detail + core
	return Room{
		PK:     pk,
		SK:     sk,
		Entity: dynamo_util.RoomEntity,
		Title:  title,
		Date:   time.Now(),
	}
}

func roomInsert(ctx context.Context, c *dynamodb.Client, instance Room) error {
	return dynamo_util.PutItem(ctx, c, instance)
}

type RoomChat struct {
	PK      string
	SK      string
	Entity  dynamo_util.EntityType
	Title   string
	Message string
	Roll    dynamo_util.Roll
	Date    time.Time
}

func createRoomChat(room Room, chat Chat) RoomChat {
	time := dynamo_util.ParseDate(time.Now())
	return RoomChat{
		PK:      room.PK,
		SK:      RPREFIX + chat.PK + "#",
		Entity:  dynamo_util.ChatEntity,
		Title:   room.Title,
		Message: chat.Message,
		Roll:    chat.Roll,
		Date:    time,
	}
}

func roomChatInsert(ctx context.Context, c *dynamodb.Client, room Room, chat Chat) error {
	roomChat := createRoomChat(room, chat)
	return dynamo_util.PutItem(ctx, c, roomChat)
}

func RoomInsert(ctx context.Context, c *dynamodb.Client, userPK string, title string) (Room, error) {
	room := roomCreate(title)
	if err := roomInsert(ctx, c, room); err != nil {
		log.Printf("Fail to create room detail %+v", err)
		return room, err
	}
	if err := userRoomInsert(ctx, c, room, userPK); err != nil {
		log.Printf("Fail to create user room %+v", err)
		return room, err
	}
	return room, nil
}
