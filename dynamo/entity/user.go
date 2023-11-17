package entity

import (
	"context"
	"dynamo/dynamo_util"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	UPrefix = "User#"
	Detail  = "Detail#"
)

type User struct {
	PK     string
	SK     string
	Entity dynamo_util.EntityType
	Name   string
}

func createUser() User {
	core := dynamo_util.IdGenerator()
	pk := UPrefix + core
	sk := UPrefix + Detail + core
	return User{
		PK:     pk,
		SK:     sk,
		Entity: dynamo_util.UserEntity,
		Name:   "Guiwoo",
	}
}

type UserRoom struct {
	PK        string
	SK        string
	Entity    dynamo_util.EntityType
	RoomTitle string
}

func createUserRoom(userPK string, room Room) UserRoom {
	return UserRoom{
		PK:        userPK,
		SK:        UPrefix + room.PK,
		Entity:    dynamo_util.RoomEntity,
		RoomTitle: room.Title,
	}
}

func userRoomInsert(ctx context.Context, c *dynamodb.Client, room Room, userPK string) error {
	userRoom := createUserRoom(userPK, room)
	return dynamo_util.PutItem(ctx, c, userRoom)
}

func UserInsert(ctx context.Context, c *dynamodb.Client) (User, error) {
	user := createUser()
	err := dynamo_util.PutItem(ctx, c, user)
	return user, err
}
