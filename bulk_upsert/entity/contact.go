package entity

type Contact struct {
	Id         string
	Title      string
	FirstName  string
	MiddleName string
	LastName   string
	Suffix     string
	Role       string
	Email      string
	Phone      string
	Mobile     string
	AltPhone   string
	Fax        string
}

func (_ Contact) TableName() string {
	return "contact"
}
