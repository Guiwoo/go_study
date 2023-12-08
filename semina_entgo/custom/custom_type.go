package custom

import (
	"database/sql/driver"
	"entgo.io/ent/schema/field"
)

type Shape string

const (
	Rectangle Shape = "정사각형"
	Square    Shape = "사각형"
	Triangle  Shape = "삼각형"
)

func (s Shape) Values() []string {
	var rs []string
	for _, v := range []Shape{Rectangle, Square, Triangle} {
		rs = append(rs, string(v))
	}
	return rs
}

type Level int

const (
	Low Level = iota
	Middle
	High
)

func (l Level) String() string {
	switch l {
	case Low:
		return "Low"
	case Middle:
		return "Middle"
	default:
		return "High"
	}
}

func (l Level) Values() []string {
	return []string{Low.String(), Middle.String(), High.String()}
}

func (l Level) Value() (driver.Value, error) {
	return l.String(), nil
}

func (l *Level) Scan(src any) error {
	var a string
	switch v := src.(type) {
	case nil:
		return nil
	case string:
		a = v
	case []uint8:
		a = string(v)
	}

	switch a {
	case "Low":
		*l = Low
	case "High":
		*l = High
	default:
		*l = Middle
	}
	return nil
}

var _ field.ValueScanner = (*Level)(nil)
