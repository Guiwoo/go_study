package dynamo_util

type EntityType int

const TableName = "tb_single_gpt"

const (
	UserEntity EntityType = iota + 1
	RoomEntity
	ChatEntity
)

type Roll string

const (
	UserRoll Roll = "user"
	GptRoll  Roll = "gpt"
)
