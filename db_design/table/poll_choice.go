package table

import (
	"math/rand"
	"time"
)

type PollChoice struct {
	Id         string    `gorm:"column:poll_choice_id;type:varchar(36);primaryKey"`
	QuestionId string    `gorm:"column:poll_question_id;type:varchar(36);index:idx_poll_question_id"`
	Choice     string    `gorm:"column:contents;type:text"`
	ImgURL     string    `gorm:"column:img_name;type:varchar(255)"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(3)"`
	Order      int       `gorm:"column:order;type:tinyint"`
	ViewType   View      `gorm:"column:view_type;type:tinyint"`
}

func (p *PollChoice) TableName() string {
	return "poll_choice"
}

func NewPollChoice(questionId, choice string) *PollChoice {
	return &PollChoice{
		Id:         Generator(),
		QuestionId: questionId,
		Choice:     choice,
		Order:      rand.Intn(100),
		ViewType:   ViewTextImage,
	}
}
