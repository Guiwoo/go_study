package table

import "time"

type Survey struct {
	Id         int       `gorm:"column:survey_id;type:int;primaryKey;autoIncrement"`
	ContentsId int       `gorm:"column:contents_id"`
	CreatorId  int       `gorm:"column:creator"`
	Title      string    `gorm:"column:title;type:varchar(20)"`
	CreatedAt  time.Time `gorm:"column:crated_at;type:datetime(3)"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime(3)"`
	Contents   Contents  `gorm:"foreignKey:contents_id"`
	Creator    User      `gorm:"foreignKey:user_id;reference:creator_id"`
}

func (s *Survey) TableName() string {
	return "survey"
}
