package singleton

import "fmt"

/**
전통적인 방식의 싱글턴
*/

type MyDB interface {
	DoSomething() // DB 작업을 위한 메서드
}

type db struct {
	// db에 해당하는 값이 저장
}

var obj *db

func init() {
	obj = &db{}
	fmt.Println("DB created")
}

func (d *db) DoSomething() {
	// DB 작업을 수행하는 메서드
}

func GetInstance() MyDB {
	if obj == nil {
		fmt.Println("There is no DB instance")
	}
	return obj
}
