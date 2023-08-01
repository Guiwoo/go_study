package test

import (
	"context"
	"db_design/table"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"testing"
)

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

	c.Db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
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
}

func TestPoll_Select(t *testing.T) {
	c := table.GetDB()
	var poll table.Poll
	err := c.Db.WithContext(context.Background()).Model(&poll).
		Preload("Contents").
		Preload("Creator").
		Preload("Updater").
		Preload("Questions").
		Preload("Questions.Choices").
		Where("content_id = ? ", 1).
		Take(&poll).Error

	if err != nil {
		t.Errorf("fail test to select poll %v \n", err)
	}

	fmt.Printf("poll data %s %s %v\n", poll.Title, poll.Creator.Name, poll)
	fmt.Println("투표질문", len(poll.Questions))
	for _, v := range poll.Questions {
		fmt.Println("----투표질문---- ", v.Title, v.Type, v.Order)
		for _, r := range v.Choices {
			fmt.Println("@@투표보기@@", r.Choice, r.Order)
		}
	}
}
