package database

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func camelToUnderscore(s string) string {
	if s == "ID" {
		return "id"
	}
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

type Mysql struct {
	db *sql.DB
}

func (m *Mysql) AutoMigrate(arr ...TableLabeler) error {
	for _, table := range arr {
		sb := strings.Builder{}
		sb.WriteString("CREATE TABLE IF NOT EXISTS " + table.TableName())
		totalFields := reflect.TypeOf(table).NumField()

		sb.WriteString("(")
		for i := 0; i < totalFields; i++ {
			field := reflect.TypeOf(table).Field(i)
			fieldName := field.Name
			fieldType := field.Type

			sb.WriteString(camelToUnderscore(fieldName))
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
		if _, err := m.db.Exec(sb.String()); err != nil {
			return err
		}
	}
	return nil
}

func (m *Mysql) BulkUpsert(arr []TableLabeler) error {
	if len(arr) == 0 {
		return nil
	}

	fields := reflect.TypeOf(arr[0]).NumField()
	columns := ""
	questions := ""
	updates := ""
	for i := 0; i < fields; i++ {
		column := camelToUnderscore(reflect.TypeOf(arr[0]).Field(i).Name)
		columns += column
		questions += "?"
		if i > 0 {
			updates += fmt.Sprintf("%s = VALUES(%s)", column, column)
			if i != fields-1 {
				updates += ","
			}
		}
		if i != fields-1 {
			columns += ","
			questions += ","
		}
	}

	query := "INSERT INTO " + arr[0].TableName() + "(" + columns + ") VALUES %s ON DUPLICATE KEY UPDATE " + updates

	var (
		value     []string
		valueArgs []interface{}
	)

	for _, v := range arr {
		value = append(value, fmt.Sprintf("(%s)", questions))
		for i := 0; i < fields; i++ {
			if reflect.TypeOf(v).Field(i).Type.Name() == "bool" {
				var val = 0
				if reflect.ValueOf(v).Field(i).Bool() {
					val = 1
				}
				valueArgs = append(valueArgs, val)
				continue
			}
			valueArgs = append(valueArgs, reflect.ValueOf(v).Field(i).String())
		}
	}

	prep := fmt.Sprintf(query, strings.Join(value, ","))

	_, err := m.db.Exec(prep, valueArgs...)

	return err
}

func (m *Mysql) Get() *sql.DB {
	return m.db
}

var _ Database = (*Mysql)(nil)

func NewMysql(dsn string) Database {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &Mysql{
		db: db,
	}
}
