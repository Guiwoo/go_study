package main

import (
	"context"
	"dynamo/dynamo_util"
	"dynamo/entity"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	client, err := newClient("")
	if err != nil {
		panic(err)
	}

	//if table, err := createTable(ctx, client); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(table)
	//}

	user := entity.User{
		PK: "User#3135fdc6-8528-11ee-b24d-2669912982eb",
	}

	room, err := entity.RoomInsert(ctx, client, user.PK, "뉴진스의 뮤직비디오 노래들의 모든 조회수에 대해 조사해줘")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 50; i++ {
		fmt.Println(i + 1)
		s := fmt.Sprintf("%d 번째 물어보기 입니다.", i+1)
		if err := entity.ChatInsert(ctx, client, room, dynamo_util.UserRoll, s); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
