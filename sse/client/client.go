package client

import (
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

var arr []string

func init() {
	n := int('z'-'a') + 1
	arr = make([]string, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, string('a'+i))
	}
	fmt.Println(arr)
}

type HubClient struct {
	log    zerolog.Logger
	alarms map[string]bool
}

func (client *HubClient) setAlarm() {
	time.Sleep(500 * time.Millisecond)
	for _, id := range arr {
		x := time.Now().Unix()
		if x%2 == 0 {
			client.alarms[id] = true
		} else {
			client.alarms[id] = false
		}
	}
}

func (client *HubClient) getAlarm(id string) bool {
	return client.alarms[id]
}

func (client *HubClient) Run() {
	defer func() {
		if err := recover(); err != nil {
			client.log.Err(fmt.Errorf("panic and recover here")).Msgf("panic and recover %+v", err)
		}
	}()

	var (
		getDbTicker = time.NewTicker(500 * time.Millisecond)
	)

	for {
		select {
		case <-getDbTicker.C:
			client.setAlarm()
		}
	}
}

func NewHubClient(log *zerolog.Logger) *HubClient {
	return &HubClient{
		log: log.With().Str("component", "HUB CLIENT").Logger(),
	}
}
