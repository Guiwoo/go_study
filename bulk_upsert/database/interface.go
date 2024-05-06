package database

import "database/sql"

type TableLabeler interface {
	TableName() string
}

type Database interface {
	AutoMigrate(arr ...TableLabeler) error
	Get() *sql.DB
	BulkUpsert(arr []TableLabeler) error
}
