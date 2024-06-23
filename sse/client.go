package main

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

type sseClient struct {
	userId string
	data   chan bool
}

type HubClient struct {
	log      zerolog.Logger
	mu       sync.Mutex
	users    map[string]*sseClient
	alarms   map[string]bool
	register chan *sseClient
}

func (client *HubClient) Register(userId string) {
	_, ok := client.users[userId]
	if ok == false {
		client.mu.Lock()
		sseClient := &sseClient{
			userId: userId,
			data:   make(chan bool),
		}
		client.users[userId] = sseClient
		client.mu.Unlock()
	}

	client.register <- client.users[userId]
}

func (client *HubClient) Unregister(userId string) {
	user, ok := client.users[userId]
	if ok {
		client.mu.Lock()
		delete(client.users, user.userId)
		client.mu.Unlock()
	}
	client.log.Debug().Msgf("%+v user clinet remove", userId)
}

func (client *HubClient) GetData(userId string) <-chan bool {
	defer client.mu.Unlock()
	client.mu.Lock()
	_, ok := client.users[userId]
	if ok {
		return client.users[userId].data
	}
	return nil
}

func (client *HubClient) getAlarm(id string) bool {
	defer client.mu.Unlock()
	client.mu.Lock()
	_, ok := client.alarms[id]
	if ok == true {
		return true
	}
	return false
}

func (client *HubClient) SetAlarm() error {
	client.mu.Lock()
	time.Sleep(1 * time.Second)
	clear(client.alarms)
	for _, id := range arr {
		x := time.Now().Unix()
		if x%2 == 0 {
			client.alarms[id] = true
		} else {
			client.alarms[id] = false
		}
	}
	client.mu.Unlock()
	return nil
}

func (client *HubClient) Run() {
	defer func() {
		if err := recover(); err != nil {
			client.log.Err(fmt.Errorf("panic and recover here")).Msgf("panic and recover %+v", err)
		}
	}()
	var (
		failCnt   = 0
		getAlarm  = time.NewTicker(3 * time.Second)
		sendAlarm = time.NewTicker(1 * time.Second)
	)

	for {
		select {
		case _ = <-getAlarm.C:
			if err := client.SetAlarm(); err != nil {
				failCnt++
				client.log.Err(err).Msgf("fail to set alarm fail count %d", failCnt)
			}
		case _ = <-sendAlarm.C:
			for _, user := range client.users {
				select {
				case user.data <- client.getAlarm(user.userId):
				default:
					client.log.Debug().Msgf("user %s channel closed or blocked", user.userId)
				}
			}
		case user := <-client.register:
			client.log.Debug().Msgf("register user %v", user.userId)
			select {
			case user.data <- client.getAlarm(user.userId):
			default:
				client.log.Debug().Msgf("user %s channel closed or blocked", user.userId)
			}
		}
	}
}

func NewHubClient(log *zerolog.Logger) *HubClient {
	return &HubClient{
		log:      log.With().Str("component", "HUB CLIENT").Logger(),
		mu:       sync.Mutex{},
		users:    make(map[string]*sseClient),
		alarms:   make(map[string]bool),
		register: make(chan *sseClient),
	}
}
