package middleware

import (
	"cache/memory_db"
	"cache/mysql_db"
)

type DataFacade interface {
	CreateData(value interface{}) error
	FindData(value interface{}) error
}

type MyServiceFacade struct {
	mysql_db.MySqlDB
	memory_db.CustomDB
}

func (m MyServiceFacade) CreateData(value interface{}) error {
	m.Create(value)
	m.Insert("name", value)
	return nil
}

func (m MyServiceFacade) FindData(value interface{}) error {
	err := m.Find(value)
	if err != nil {
		_, err = m.Get("name")
		if err != nil {
			return err
		}
	}
	return nil
}

var _ DataFacade = (*MyServiceFacade)(nil)
