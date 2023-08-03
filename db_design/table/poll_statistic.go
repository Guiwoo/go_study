package table

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PollStatistic struct {
	ChoiceId   string         `gorm:"column:poll_choice_id;type:varchar(36);"`
	PollId     string         `gorm:"column:poll_id;type:varchar(36);"`
	QuestionId string         `gorm:"column:poll_question_id;type:varchar(36)"`
	Polled     int            `gorm:"column:polled;type:bigint(20)"`
	Poll       Poll           `gorm:"foreignKey:poll_id;reference:PollId"`
	Result     PollResultData `gorm:"foreignKey:poll_choice_id;reference:poll_choice_id"`
}

func (p *PollStatistic) TableName() string {
	return "poll_statistic"
}

func (p *PollStatistic) Select(db *gorm.DB) (result []PollStatistic, err error) {
	err = db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
		var ids []string
		if err = tx.Model(&Poll{}).Select("poll_id").Where("status = ?", 1).Find(&ids).Error; err != nil {
			return err
		}
		if err = tx.Model(&PollStatistic{}).Where("poll_id in (?)", ids).Find(&result).Error; err != nil {
			return err
		}
		return nil
	})
	return result, err
}

func (p *PollStatistic) Select2(db *gorm.DB) (result []PollStatistic, err error) {
	err = db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
		subquery := tx.Model(&Poll{}).Select("poll_id").Where("status = ? ", PollProgress)
		if err = tx.Model(&PollStatistic{}).
			Preload("Result").
			Preload("Poll").
			Preload("Poll.Questions").
			Where("poll_id in (?)", subquery).
			Group("poll_id").Group("poll_question_id").Group("poll_choice_id").
			Find(&result).Error; err != nil {
			return err
		}
		return nil
	})
	return result, err
}

func (p *PollStatistic) Upsert(db *gorm.DB) error {
	return db.WithContext(context.Background()).Transaction(func(tx *gorm.DB) error {
		// 진행중인 poll 가져와서 넣어주기
		var list []Poll
		if err := tx.Model(&Poll{}).Find(&list).Where("status = ? ", PollProgress).Error; err != nil {
			return err
		}
		//인서트 떄리자 duplicate 무시해서 떄리면 될듯 ?
		var rsList []PollStatistic
		if err := tx.Model(&PollResult{}).
			Select("poll_id", "poll_question_id", "poll_choice_id", "count(poll_id) as polled").
			Group("poll_choice_id").
			Group("poll_question_id").
			Group("poll_id").
			Find(&rsList).
			Error; err != nil {
			return err
		}

		for _, v := range rsList {
			if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&v).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
