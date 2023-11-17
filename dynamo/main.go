package main

import (
	"context"
	"dynamo/dynamo_util"
	"dynamo/entity"
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

	user, err := entity.UserInsert(ctx, client)
	if err != nil {
		panic(err)
	}

	room, err := entity.RoomInsert(ctx, client, user.PK, "Intel Cpu 의 하이퍼 스레딩이 뭐야?")
	if err != nil {
		panic(err)
	}

	if err := entity.ChatInsert(ctx, client, room, dynamo_util.UserRoll, "Intel Cpu의 하이퍼 스레딩이 뭐야?"); err != nil {
		panic(err)
	}
}
