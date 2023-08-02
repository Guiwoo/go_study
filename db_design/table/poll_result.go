package table

import (
	"math/rand"
	"strings"
	"time"
)

type PollResult struct {
	PollResultId string    `gorm:"column:poll_result_id;type:varchar(36);index:idx_poll_result_id"`
	UserId       int       `gorm:"column:user_id;type:bigint(20);index:idx_poll_result_id"`
	PollId       string    `gorm:"column:poll_id;type:varchar(36);index:idx_poll_id"`
	QuestionId   string    `gorm:"column:poll_question_id;type:varchar(36);index:idx_question_id"`
	PollChoiceId string    `gorm:"column:poll_choice_id;type:varchar(36);index:idx_poll_choice_id"`
	PollResult   string    `gorm:"column:poll_result;type:text"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(3)"`
}

func (p *PollResult) TableName() string {
	return "poll_result"
}

func getRs(c PollChoice) string {
	if strings.TrimSpace(c.Choice) == "" {
		return "서술형 답변"
	}
	return c.Choice
}

func NewPollResult(p Poll, cnt int) *PollResult {
	ran := rand.Intn(100000)
	ques := p.Questions[ran%len(p.Questions)]
	choice := ques.Choices[ran%len(ques.Choices)]
	return &PollResult{
		PollResultId: Generator(),
		PollId:       p.Id,
		UserId:       ran%cnt + 1,
		QuestionId:   ques.Id,
		PollChoiceId: choice.Id,
		PollResult:   getRs(choice),
	}
}
