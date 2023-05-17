package flyweight

import (
	"fmt"
	"strings"
	"unicode"
)

/**
Avoid redundancy when storing data
Ex)
- Plenty of users with identical first/last name
- No sense in storing same first/last name over and over again
- Store a list of names and references to them


A space optimization technique that lets us use less memory by storing externally the data associated with similar objects=
*/

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{plainText, make([]bool, len(plainText))}
}

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r
}

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

type User2 struct {
	names []uint8
}

func NewUser2(fullName string) *User2 {
	getOrdAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrdAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func Start() {
	text := "There are three pigs in the farm"

	ft := NewFormattedText(text)
	ft.Capitalize(10, 15)
	fmt.Println(ft.String())
}
