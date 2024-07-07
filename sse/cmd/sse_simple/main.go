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

var alarms map[string]bool

func init() {
	a := sync.Once{}
	a.Do(func() {
		if alarms == nil {
			alarms = make(map[string]bool)
			for i := 'a'; i <= 'z'; i++ {
				alarms[string(i)] = true
			}
		}
	})
}

func getData(id string) bool {
	// 데이터 조회해오는 로직
	for i := 'a'; i <= 'z'; i++ {
		x := time.Now().Unix()
		if x%2 == 0 {
			alarms[string(i)] = true
		} else {
			alarms[string(i)] = false
		}
	}

	return alarms[id]
}

func main() {
	e := echo.New()
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	e.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{}),
	)

	sse := e.Group("/sse")

	sse.GET("/call/:id", func(c echo.Context) error {
		time.Sleep(500 * time.Millisecond)
		id := c.Param("id")
		done := make(chan bool)
		log.Info().Msgf("Get Request Id :%+v", id)
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
					str := fmt.Sprintf(`{"alarm":%v}`, getData(id))
					if _, err := fmt.Fprintf(c.Response().Writer, "data: %s\n\n", str); err != nil {
						log.Err(err).Msg("failed to send data")
					}

					c.Response().Flush()
				case <-c.Request().Context().Done():
					log.Info().Msg("get done signal from request")
					done <- true
					return
				}
			}
		}(done, id)
		<-done

		log.Info().Msgf("pass the groutine")
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
