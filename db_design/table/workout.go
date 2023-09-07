package table

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type Workout struct {
	Id    int    `gorm:"column:id;type:int;primaryKey"`
	Title string `gorm:"column:title;type:varchar(20)"`
	Raps  int    `gorm:"column:raps;type:int"`
}

func (w *Workout) TableName() string {
	return "workout"
}

func (w *Workout) Insert(db *gorm.DB) error {
	return db.Create(w).Error
}

func (w *Workout) Upsert(db *gorm.DB) error {
	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"raps": gorm.Expr("raps + ?", w.Raps),
		}),
	}).Create(w).Error
}

func (w *Workout) BatchUpsert(db *gorm.DB, data []Workout) error {
	var (
		value     []string
		valueArgs []interface{}
	)

	for _, v := range data {
		value = append(value, ("(?,?,?)"))

		valueArgs = append(valueArgs, v.Id)
		valueArgs = append(valueArgs, v.Title)
		valueArgs = append(valueArgs, v.Raps)
	}

	prep := "insert into workout(id,title,raps) values %s on duplicate key update raps = raps+values(raps)"

	sql := fmt.Sprintf(prep, strings.Join(value, ","))

	fmt.Println(sql)

	if err := db.Exec(sql, valueArgs...).Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

func (w *Workout) BatchUpdate(db *gorm.DB, data []Workout) error {
	var (
		caseSql   []string
		whereSql  []string
		caseArgs  []interface{}
		whereArgs []interface{}
	)

	for _, v := range data {
		caseSql = append(caseSql, "when ? then raps + ?")
		caseArgs = append(caseArgs, v.Id, v.Raps)
		whereArgs = append(whereArgs, v.Id)
		whereSql = append(whereSql, "?")
	}

	prep := "update workout set raps = case id %s end where id in (%s)"

	sql := fmt.Sprintf(prep, strings.Join(caseSql, " "), strings.Join(whereSql, ","))

	caseArgs = append(caseArgs, whereArgs...)

	if err := db.Exec(sql, caseArgs...).Error; err != nil {
		return err
	}
	return nil
}
