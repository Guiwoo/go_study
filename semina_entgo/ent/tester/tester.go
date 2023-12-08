// Code generated by ent, DO NOT EDIT.

package tester

import (
	"fmt"
	"semina_entgo/custom"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tester type in the database.
	Label = "tester"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPascalCase holds the string denoting the pascalcase field in the database.
	FieldPascalCase = "pascal_case"
	// FieldLetMeCheck holds the string denoting the let_me_check field in the database.
	FieldLetMeCheck = "let_me_check"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldShape holds the string denoting the shape field in the database.
	FieldShape = "shape"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// Table holds the table name of the tester in the database.
	Table = "testers"
)

// Columns holds all SQL columns for tester fields.
var Columns = []string{
	FieldID,
	FieldPascalCase,
	FieldLetMeCheck,
	FieldSize,
	FieldShape,
	FieldLevel,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Size defines the type for the "size" enum field.
type Size string

// Size values.
const (
	SizeBig   Size = "big"
	SizeSmall Size = "small"
)

func (s Size) String() string {
	return string(s)
}

// SizeValidator is a validator for the "size" field enum values. It is called by the builders before save.
func SizeValidator(s Size) error {
	switch s {
	case SizeBig, SizeSmall:
		return nil
	default:
		return fmt.Errorf("tester: invalid enum value for size field: %q", s)
	}
}

// ShapeValidator is a validator for the "shape" field enum values. It is called by the builders before save.
func ShapeValidator(s custom.Shape) error {
	switch s {
	case "정사각형", "사각형", "삼각형":
		return nil
	default:
		return fmt.Errorf("tester: invalid enum value for shape field: %q", s)
	}
}

// LevelValidator is a validator for the "level" field enum values. It is called by the builders before save.
func LevelValidator(l custom.Level) error {
	switch l.String() {
	case "Low", "Middle", "High":
		return nil
	default:
		return fmt.Errorf("tester: invalid enum value for level field: %q", l)
	}
}

// OrderOption defines the ordering options for the Tester queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPascalCase orders the results by the PascalCase field.
func ByPascalCase(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPascalCase, opts...).ToFunc()
}

// ByLetMeCheck orders the results by the let_me_check field.
func ByLetMeCheck(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLetMeCheck, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByShape orders the results by the shape field.
func ByShape(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShape, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}