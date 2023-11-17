package entity

import (
	"context"
	"dynamo/dynamo_util"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

const (
	CPrefix = "Chat#"
)

type Chat struct {
	PK        string
	SK        string
	Entity    dynamo_util.EntityType
	Roll      dynamo_util.Roll
	Message   string
	CreatedAt time.Time
}

func createChat(roll dynamo_util.Roll, message string) Chat {
	core := dynamo_util.IdGenerator()
	pk := CPrefix + core
	sk := CPrefix + Detail + core
	return Chat{
		PK:        pk,
		SK:        sk,
		Entity:    dynamo_util.ChatEntity,
		Roll:      roll,
		Message:   message,
		CreatedAt: time.Now(),
	}
}

func ChatInsert(ctx context.Context, c *dynamodb.Client, room Room, roll dynamo_util.Roll, message string) error {
	chat := createChat(roll, message)
	if err := dynamo_util.PutItem(ctx, c, chat); err != nil {
		return err
	}
	if err := roomChatInsert(ctx, c, room, chat); err != nil {
		return err
	}
	return nil
}
