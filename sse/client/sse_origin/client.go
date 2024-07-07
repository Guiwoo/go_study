package sse_origin

import (
	"fmt"
	"github.com/rs/zerolog"
	"sync"
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
	mutex  sync.Mutex
}

func (client *HubClient) setAlarm() {
	time.Sleep(500 * time.Millisecond)
	client.mutex.Lock()
	for _, id := range arr {
		x := time.Now().Unix()
		if x%2 == 0 {
			client.alarms[id] = true
		} else {
			client.alarms[id] = false
		}
	}
	client.mutex.Unlock()
}

func (client *HubClient) GetAlarm(id string) bool {
	a := false
	client.mutex.Lock()
	a = client.alarms[id]
	client.mutex.Unlock()
	return a
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
		log:    log.With().Str("component", "HUB CLIENT").Logger(),
		alarms: make(map[string]bool),
		mutex:  sync.Mutex{},
	}
}
