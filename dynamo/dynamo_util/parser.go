package dynamo_util

import (
	"log"
	"time"
)

func ParseDate(t time.Time) time.Time {
	parsed := t.Format("2006-01-02T15:04:05")
	loc, _ := time.LoadLocation("Asia/Seoul")
	newTime, err := time.ParseInLocation("2006-01-02T15:04:05", parsed, loc)
	if err != nil {
		log.Printf("parse time error %+v", err)
	}
	return newTime
}
