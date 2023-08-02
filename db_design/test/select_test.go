package test

import (
	"context"
	"db_design/table"
	"fmt"
	"testing"
)

func TestUserSeed(t *testing.T) {

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

func TestPollStatistic_Select(t *testing.T) {
	db := table.GetDB()
	ps := table.PollStatistic{}
	list, err := ps.Select(db.Db)
	if err != nil {
		t.Errorf("select error on poll statistic %v", err)
	}
	fmt.Println("✅")
	fmt.Println(list)
}
