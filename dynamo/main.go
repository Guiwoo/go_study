package main

import (
	"context"
	"dynamo/dynamo_util"
	"dynamo/entity"
	"fmt"
	"os"
	"time"
)

func main() {
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")
	fmt.Println(key, secret)

	ctx := context.Background()
	client, err := newClient(key, secret)
	if err != nil {
		panic(err)
	}

	if table, err := createTable(ctx, client); err != nil {
		panic(err)
	} else {
		fmt.Println(table)
	}
	fmt.Println("create table")

	user, err := entity.UserInsert(ctx, client)
	if err != nil {
		panic(err)
	}
	fmt.Println("create user")

	room, err := entity.RoomInsert(ctx, client, user.PK, "뉴진스의 뮤직비디오 노래들의 모든 조회수에 대해 조사해줘")
	if err != nil {
		panic(err)
	}

	fmt.Println("create room")

	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
		s := fmt.Sprintf("%d 번째 물어보기 입니다.", i+1)
		if err := entity.ChatInsert(ctx, client, room, dynamo_util.UserRoll, s); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
