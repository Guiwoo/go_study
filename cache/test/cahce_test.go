package test

import (
	"cache/memory_db"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func TestCacheInsertAndGet(t *testing.T) {
	memory_db.Init(5, 1)

	db := memory_db.GetCacheDB()

	key := "person"

	db.Insert(key, struct {
		name string
		age  int
	}{
		"guiwoo", 30,
	})

	data, err := db.Get(key)

	if err != nil || data == nil {
		t.Error(err)
	}

	fmt.Println(data)

	logrus.Info("time will sleep for 5 seconds")
	time.Sleep(5 * time.Second)

	data, err = db.Get(key)

	if err == nil || data != nil {
		t.Error(err)
	}

	fmt.Println(data)
}

func TestCheckValueIsPointer(t *testing.T) {
	type test struct {
		name, email string
	}

	a := test{"a", "a@a.com"}
	b := &test{"b", "b@b.com"}

	isPointer := func(value any) bool {
		if reflect.ValueOf(value).Kind() == reflect.Pointer {
			return true
		}
		return false
	}

	fmt.Printf("is pointer a : %v\n", isPointer(a))
	fmt.Printf("is pointer b : %v\n", isPointer(b))
}

func TestCoverExistValue(t *testing.T) {
	memory_db.Init(5, 10)
	db := memory_db.GetCacheDB()

	type person struct {
		name string
	}

	err := db.Insert("person", &person{"guiwoo"})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(4 * time.Second)
	err = db.Insert("person", &person{"changed"})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(4 * time.Second)

	data, err := db.Get("person")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data.(*person))
}
