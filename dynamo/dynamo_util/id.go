package dynamo_util

import "github.com/google/uuid"

func IdGenerator() string {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return u.String()
}
