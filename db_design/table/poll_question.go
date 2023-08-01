package table

import (
	"math/rand"
	"time"
)

type PollQuestion struct {
	Id        string       `gorm:"column:poll_question_id;type:varchar(36);primaryKey"`
	PollId    string       `gorm:"column:poll_id;index:idx_poll_id"`
	Title     string       `gorm:"column:title;type:varchar(1000)"`
	Type      QuestionType `gorm:"column:question_type;type:varchar(1)"`
	CreatedAt time.Time    `gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt time.Time    `gorm:"column:updated_at;type:datetime(3)"`
	ImgURL    string       `gorm:"column:img_name;type:varchar(255)"`
	Order     int          `gorm:"column:order;type:tinyint"`
	IsUse     Using        `gorm:"column:is_use;type:tinyint"`
	ViewType  View         `gorm:"column:view_type;type:tinyint"`
	Choices   []PollChoice `gorm:"foreignKey:poll_question_id;reference:poll_question_id"`
}

func (p *PollQuestion) TableName() string {
	return "poll_question"
}

func NewPollQuestion(pollId, title string, t QuestionType) *PollQuestion {
	return &PollQuestion{
		Id:     Generator(),
		PollId: pollId,
		Title:  title,
		Type:   t,
		Order:  rand.Intn(100),
	}
}
