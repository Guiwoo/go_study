package faker

import "github.com/brianvoe/gofakeit/v6"

type FakeUser struct {
	Name string `fake:"{firstname}"`
}

func (f *FakeUser) TableName() string {
	return "users"
}

func NewFakeUser() *FakeUser {
	var f FakeUser
	if err := gofakeit.Struct(&f); err != nil {
		return nil
	}
	return &f
}
