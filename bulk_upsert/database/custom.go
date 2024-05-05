package database

import (
	"bytes"
	"database/sql"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func camelToUnderscore(s string) string {
	var buf bytes.Buffer

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

type Custom struct {
	db *sql.DB
}

func (c *Custom) AutoMigrate(arr ...TableLabeler) error {
	for _, table := range arr {
		sb := strings.Builder{}
		sb.WriteString("CREATE TABLE IF NOT EXISTS " + table.TableName())
		totalFields := reflect.TypeOf(table).NumField()

		sb.WriteString("(")
		for i := 0; i < totalFields; i++ {
			field := reflect.TypeOf(table).Field(i)
			fieldName := field.Name
			fieldType := field.Type

			sb.WriteString(fieldName)
			if fieldType.Name() == "string" {
				sb.WriteString(" varchar(255) ")
			}
			if fieldType.Name() == "bool" {
				sb.WriteString(" tinyint(1) ")
			}

			if i == 0 {
				sb.WriteString("NOT NULL PRIMARY KEY")
			}

			if i < totalFields-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString(")")
		if _, err := c.db.Exec(sb.String()); err != nil {
			return err
		}
	}
	return nil
}

func NewCustom() *Custom {
	db, err := sql.Open("mysql", "root:rain45bow@/test")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Custom{
		db: db,
	}
}
