package test

import (
	"db_design/table"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestScenario_01(t *testing.T) {
	db := table.GetDB()
	p := table.Poll{}
	if err := p.Select(db.Db); err != nil {
		t.Errorf("poll select error %v", err)
	}
	cnt, err := db.Count(&table.User{})
	if err != nil {
		t.Errorf("user count error %v", err)
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	var sc sync.WaitGroup

	sc.Add(1)
	go func() {
		defer func() {
			sc.Done()
		}()
	Loop:
		for {
			select {
			case _, ok := <-done:
				if !ok {
					break Loop
				}
			default:
				//fmt.Println("대기")
				pr := table.NewPollResult(p, int(cnt))
				err = db.Insert(pr)
				if err != nil {
					fmt.Printf("err %v\n", err)
				}
				err = db.Update("participants", &table.Poll{}, p.Id)
				if err != nil {
					fmt.Printf("update err %v\n", err)
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	tc := time.NewTicker(800 * time.Millisecond)
	sc.Add(1)
	go func() {
		defer sc.Done()
	Loop:
		for {
			select {
			case _, ok := <-done:
				if !ok {
					break Loop
				}
			case <-tc.C:
				ps := table.PollStatistic{}
				ps.Upsert(db.Db)
			}
		}
	}()
	//
	tc2 := time.NewTicker(1 * time.Second)
	sc.Add(1)
	go func() {
		defer sc.Done()
	Loop:
		for {
			select {
			case _, ok := <-done:
				if !ok {
					break Loop
				}
			case <-tc2.C:
				// 조회 1초에 한번
				ps := table.PollStatistic{}
				result, _ := ps.Select(db.Db)
				for _, v := range result {
					fmt.Println(v.PollId)
				}
			}
		}
	}()

	sc.Wait()
}
