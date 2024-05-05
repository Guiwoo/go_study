package entity

type Contract struct {
	ID         string
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

func (_ Contract) TableName() string {
	return "contract"
}
