package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sse/client/third"
	"time"
)

func main() {
	e := echo.New()
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	hub := third.NewHubClient(&log)

	go hub.Run()

	e.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{}),
	)

	done := make(chan bool)
	sse := e.Group("/sse")
	sse.GET("/call/:id", func(c echo.Context) error {
		id := c.Param("id")
		log.Info().Msgf("Get Request Id :%+v", id)
		hub.Connect(id)
		log.Info().Msgf("after register")
		go func(done chan bool, id string) {
			ticker := time.NewTicker(500 * time.Millisecond)
			defer log.Info().Msgf("go routine done")
			for {
				select {
				case <-ticker.C:
					c.Response().Header().Set("Content-Type", "text/event-stream")
					c.Response().Header().Set("Cache-Control", "no-cache")
					c.Response().Header().Set("Connection", "keep-alive")
					str := fmt.Sprintf(`{"alarm":%v}`, hub.GetAlarm(id))
					if _, err := fmt.Fprintf(c.Response().Writer, "data: %s\n\n", str); err != nil {
						log.Err(err).Msg("failed to send data")
					}

					c.Response().Flush()
				case <-c.Request().Context().Done():
					log.Info().Msg("get done signal from request")
					hub.Disconnect(id)
					done <- true
					return
				}
			}
		}(done, id)
		<-done

		log.Info().Msgf("pass the groutine")
		log.Debug().Msgf("sse conection has been closed")

		log.Debug().Msgf("sse conection has been closed")
		return nil
	})

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	if err := e.Start(":8080"); err != nil {
		log.Err(err).Msg("fail to start server")
	}
}
