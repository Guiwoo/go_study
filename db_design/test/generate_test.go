package test

import (
	"context"
	"db_design/faker"
	"db_design/table"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"testing"
)

func insertFakeUser(needs int) error {
	db := table.GetDB()

	for i := 0; i < needs; i++ {
		fakeUser := faker.NewFakeUser()
		if err := db.Insert(fakeUser); err != nil {
			return fmt.Errorf("%d occured error : %v", i+1, err)
		}
	}
	return nil
}

func Test_FakeUser(t *testing.T) {
	db := table.GetDB()

	total, err := db.Count(&table.User{})
	if err != nil {
		t.Errorf("count user error %v", err)
	} else if total < 10000 {
		t.Errorf("total user less than 10000 need to run insertFakeUser %d needs %d", total, 10000-total)
	}
}

func TestUpdatePoll(t *testing.T) {
	db := table.GetDB()
	id := "c642196c-3042-11ee-baa7-2669912982ea"
	if err := db.Update("participants", table.Poll{}, id); err != nil {
		t.Errorf("update participnats error %v", err)
	}
}

func TestUUID(t *testing.T) {
	rst, _ := uuid.NewUUID()
	result := rst.String()
	fmt.Println(result, len(result))
}

func TestPoll_Insert(t *testing.T) {
	c := table.GetDB()

	err := c.Insert(table.NewPoll("TestPoll", 1, 1))
	if err != nil {
		t.Errorf("fail test to insert poll %v\n", err)
	}
}

func TestPollQuestion_Insert(t *testing.T) {
	c := table.GetDB()
	poll := "c642196c-3042-11ee-baa7-2669912982ea" // select 로 들고 와서 사용할것
	err := c.Insert(table.NewPollQuestion(poll, "투표 3번 항목", table.Description))
	if err != nil {
		t.Errorf("fail test to insert poll %v\n", err)
	}
}

func TestQuestionChoice_Insert(t *testing.T) {
	c := table.GetDB()
	q1 := "78e454ac-3045-11ee-a18e-2669912982ea"
	q2 := "82d2ff18-3045-11ee-ba69-2669912982ea"
	q3 := "9397bb30-304b-11ee-b9d4-2669912982ea"

	err := c.Db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(table.NewPollChoice(q1, "보기 1번")).Error
		if err != nil {
			return err
		}
		err = tx.Create(table.NewPollChoice(q1, "보기 2번")).Error
		if err != nil {
			return err
		}
		err = tx.Create(table.NewPollChoice(q2, "보기 1번")).Error
		if err != nil {
			return err
		}
		err = tx.Create(table.NewPollChoice(q2, "보기 2번")).Error
		if err != nil {
			return err
		}
		err = tx.Create(table.NewPollChoice(q3, "서술형")).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Errorf("question choice insert fail %v", err)
	}
}

func TestPollResult(t *testing.T) {
	db := table.GetDB()

	var p table.Poll
	if err := p.Select(db.Db); err != nil {
		t.Errorf("select error %v", err)
	}
	cnt, err := db.Count(&table.User{})
	if err != nil {
		t.Errorf("count fail %v", err)
	}

	var failCnt, successCnt int
	for i := 0; i < int(cnt); i++ {
		pr := table.NewPollResult(p, int(cnt))
		if err = db.Insert(pr); err != nil {
			log.Errorf("fail to insert %v", err)
			failCnt++
			continue
		}
		successCnt++
	}

	fmt.Printf("total %d  == %d (success %d fail %d)", cnt, successCnt+failCnt, successCnt, failCnt)
}

func TestPollStatisticUpsert(t *testing.T) {
	db := table.GetDB()
	ps := table.PollStatistic{}

	ps.Upsert(db.Db)
}
