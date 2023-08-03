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
	list, err := ps.Select2(db.Db)
	if err != nil {
		t.Errorf("select error on poll statistic %v", err)
	}
	for _, v := range list {
		fmt.Printf("투표 아이디 : %s/ 질문 아이디 : %s / 보기 아이디 : %s /답변횟수 %d /답변 : %v\n", v.PollId, v.QuestionId, v.ChoiceId, v.Polled, v.Result)
		fmt.Printf("투표 타이틀 %s / 투표 시작시간 %v / 투표 종료시간 %v / 투표참여자 %d \n", v.Poll.Title, v.Poll.StartTime, v.Poll.EndTime, v.Poll.Participants)
		for _, vv := range v.Poll.Questions {
			fmt.Printf("질문 타이틀 %s / 질문 타입 %s / 질문 순서 %v \n", vv.Title, vv.Type, vv.Order)
		}
		fmt.Println()
		//for _, vv := range v.Poll.Questions {
		//	fmt.Println("질문 : ", vv.Title, vv.Type, vv.IsUse)
		//	for _, vvv := range vv.Choices {
		//		fmt.Printf("보기 %s : 답변 %s 득표수 : %d\n", vvv.Choice, v.Result.PollResult, v.Polled)
		//	}
		//}
	}
}
