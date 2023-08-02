package test

import (
	"db_design/table"
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
1000 개 의 고루틴이 각각 result 에 데이터를 쌓고 한번의 고루틴 10 개의 result 데이터를 적재한후 2초의 대기시간을 소모한후 종료한다.
매 1초 마다. result 테이블의 결과를 조회해 statistic 테이블 을 업데이트 하고 poll 테이블은 업데이트를 한다,
투표 가 종료되는 시점 statistic 테이블 과 poll 테이블은 업데이트를 한다.
*/

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

	done := make(chan interface{}) // 신호가 가는지보
	//종료신호 약 10초 후에
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
				// 폴 아이디로 다 긁어서 넣어야 하잖아 ?
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
