package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexpCardNumber(t *testing.T) {
	rgx := regexp.MustCompile("([0-9]{4})-([0-9]{4})-([0-9]{4})-([0-9]{4})")
	fmt.Println(rgx.Match([]byte("1234-1234-1234-1234")))
}
