package entity

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"unicode"
)

func CamelToUnderscore(s string) string {
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

func TestAddress_AutoMigrate(t *testing.T) {
	a := reflect.TypeOf(Vendor{}).NumField()

	for i := 0; i < a; i++ {
		f := reflect.TypeOf(Vendor{}).Field(i)
		fmt.Println(CamelToUnderscore(f.Name), f.Type.Name())
	}
}
