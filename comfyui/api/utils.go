package api

import (
	"log"
	"os"
)

func getEnv(env string) string {
	value, ok := os.LookupEnv(env)
	if ok == false {
		log.Fatalf("fail to get env %+v", env)
		return ""
	}
	return value
}
