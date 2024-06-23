package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"
)

func main() {
	e := echo.New()
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	hub := NewHubClient(&log)

	go hub.Run()

	e.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{}),
	)

	sse := e.Group("/sse")

	sse.GET("/call/:id", func(c echo.Context) error {
		time.Sleep(500 * time.Millisecond)
		id := c.Param("id")
		log.Info().Msgf("Get Request Id :%+v", id)
		hub.Register(id)
		log.Info().Msgf("after register")
		var wg sync.WaitGroup
		wg.Add(1)
		go func(id string, ctx echo.Context, wg *sync.WaitGroup) {
			defer log.Info().Msg("go routine is done")
			defer wg.Done()
			for {
				select {
				case data := <-hub.GetData(id):
					c.Response().Header().Set("Content-Type", "text/event-stream")
					c.Response().Header().Set("Cache-Control", "no-cache")
					c.Response().Header().Set("Connection", "keep-alive")
					str := fmt.Sprintf(`{"alarm":%v}`, data)
					if _, err := fmt.Fprintf(c.Response().Writer, "data: %s\n\n", str); err != nil {
						log.Err(err).Msg("failed to send data")
					}

					c.Response().Flush()
				case <-c.Request().Context().Done():
					log.Info().Msg("get done signal from request")
					hub.Unregister(id)
					return
				}
			}
		}(id, c, &wg)
		log.Info().Msgf("pass the groutine")
		wg.Wait()

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
