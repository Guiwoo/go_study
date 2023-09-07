package test

import (
	"db_design/table"
	"testing"
)

func TestWorkoutInsertTest(t *testing.T) {
	db := table.GetDB().Db

	deadLift := &table.Workout{
		Id: 1, Title: "데드리프트", Raps: 1,
	}
	benchPress := &table.Workout{
		Id: 2, Title: "벤치프레스", Raps: 1,
	}
	squat := &table.Workout{
		Id: 3, Title: "스쿼트", Raps: 1,
	}

	err := deadLift.Insert(db)
	if err != nil {
		t.Error(err)
	}

	err = benchPress.Insert(db)
	if err != nil {
		t.Error(err)
	}

	err = squat.Insert(db)
	if err != nil {
		t.Error(err)
	}
}

func TestWorkoutUpsert(t *testing.T) {
	db := table.GetDB().Db

	for i := 1; i <= 3; i++ {
		w := &table.Workout{
			Id:   i,
			Raps: i * i * 3,
		}
		if err := w.Upsert(db); err != nil {
			t.Error(err)
		}
	}
}

func TestWorkoutBatchUpsert(t *testing.T) {
	db := table.GetDB().Db
	list := make([]table.Workout, 0, 3)

	for i := 1; i <= 3; i++ {
		w := table.Workout{
			Id:   i,
			Raps: i * i * 3,
		}
		list = append(list, w)
	}

	var workout table.Workout

	if err := workout.BatchUpsert(db, list); err != nil {
		t.Error(err)
	}
}

func TestWorkoutUpdateRaw(t *testing.T) {
	db := table.GetDB().Db

	sql := "update workout set raps = case id " +
		"when ? then raps + ? " +
		"when ? then raps + ? " +
		"when ? then raps + ?" +
		" end where id in (?,?,?)"

	args := []interface{}{1, 10, 2, 20, 3, 30, 1, 2, 3}

	err := db.Exec(sql, args...).Error

	if err != nil {
		t.Error(err)
	}
}

func TestWorkoutBatchUpdate(t *testing.T) {
	db := table.GetDB().Db

	list := make([]table.Workout, 0, 3)

	for i := 1; i <= 3; i++ {
		w := table.Workout{
			Id:   i,
			Raps: i * i * 3,
		}
		list = append(list, w)
	}

	var workout table.Workout

	err := workout.BatchUpdate(db, list)

	if err != nil {
		t.Error(err)
	}
}

func TestBatchUpsertAndUpdate(t *testing.T) {
	db := table.GetDB().Db
	list := make([]table.Workout, 0, 1000)

	for i := 1; i <= 1000; i++ {
		w := table.Workout{
			Id:   (i%3 + 1),
			Raps: i + i,
		}
		list = append(list, w)
	}

	var workout table.Workout

	workout.BatchUpsert(db, list)

}
func TestBatchUpdate(t *testing.T) {
	db := table.GetDB().Db
	list := make([]table.Workout, 0, 1000)

	for i := 1; i <= 1000; i++ {
		w := table.Workout{
			Id:   (i%3 + 1),
			Raps: i + i,
		}
		list = append(list, w)
	}

	var workout table.Workout
	workout.BatchUpdate(db, list)
}
