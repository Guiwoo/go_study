package main

import (
	"flag"
	"gin_pra/config"
	"github.com/gin-gonic/gin"
	"log"
)

var envPath = flag.String("c", "", "example) go run main.go -c env path")

func main() {
	flag.Parse()

	cfg := config.SetConfig(*envPath)

	router := gin.New()
	gin.SetMode(cfg.Mode)
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("fail to set trusted proixes %+v", err)
	}

	if err := router.Run(cfg.Host); err != nil {
		log.Fatalf("fail to run server %+v", err)
	}
}
