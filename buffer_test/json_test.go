package main

import (
	"encoding/json"
	fastjson "github.com/goccy/go-json"
	"testing"
)

type Address struct {
	City        string
	ZipCode     string
	PostCode    uint32
	CountryCode uint16
	CityCode    uint16
	People      uint8
}

func BenchmarkJsonParser(b *testing.B) {
	seoul := Address{
		"Seoul", "117128", 11731, 82, 02, 128,
	}
	byteData, _ := json.Marshal(seoul)
	b.Run("Standard Json Marshal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			json.Marshal(seoul)
		}
	})
	b.Run("Json Library Marshal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fastjson.Marshal(seoul)
		}
	})
	b.Run("Standard Json UnMarshal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			json.Unmarshal(byteData, &seoul)
		}
	})
	b.Run("Json Library UnMarshal", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fastjson.Unmarshal(byteData, &seoul)
		}
	})
}
