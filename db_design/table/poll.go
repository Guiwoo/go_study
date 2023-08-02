package table

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Poll struct {
	Id           string         `gorm:"column:poll_id;type:varchar(36);primaryKey"`
	ContentsId   int            `gorm:"column:content_id;index:idx_content_id"`
	CreatorId    int            `gorm:"column:create_id;index:idx_create_id"`
	UpdateId     int            `gorm:"column:update_id;index:idx_update_id"`
	Title        string         `gorm:"column:title;type:varchar(20)"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime(3)"`
	IsUse        Using          `gorm:"column:is_use;type:tinyint"`
	IsPrivate    int            `gorm:"column:is_private;type:tinyint"`
	StartTime    time.Time      `gorm:"column:start_time;type:datetime(3)"`
	EndTime      time.Time      `gorm:"column:end_time;type:datetime(3)"`
	Status       PollStatus     `gorm:"column:status;type:tinyint"`
	Participants int            `gorm:"column:participants;type:int"`
	Contents     Contents       `gorm:"foreignKey:content_id;reference:contents_id"`
	Creator      User           `gorm:"foreignKey:create_id;reference:create_id"`
	Updater      User           `gorm:"foreignKey:update_id;reference:update_id"`
	Questions    []PollQuestion `gorm:"foreignKey:poll_id;reference:poll_id"`
}

func (p *Poll) TableName() string {
	return "poll"
}

func (p *Poll) Select(db *gorm.DB) error {
	return db.WithContext(context.Background()).Model(p).
		Preload("Contents").
		Preload("Creator").
		Preload("Updater").
		Preload("Questions").
		Preload("Questions.Choices").
		Where("content_id = ? ", 1).
		Take(p).Error
}

func NewPoll(title string, contentsId, adminId int) *Poll {
	t := time.Now()
	return &Poll{
		Id:         Generator(),
		ContentsId: contentsId,
		CreatorId:  adminId,
		UpdateId:   adminId,
		Title:      title,
		Status:     PollProgress,
		StartTime:  t,
		EndTime:    t.Add(3 * time.Hour),
	}
}
