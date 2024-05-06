package entity

type Address struct {
	Id         string
	Person     string
	Company    string
	Line1      string
	Line2      string
	Line3      string
	City       string
	State      string
	PostalCode string
	Country    string
	Email      string
	Phone      string
}

func (_ Address) TableName() string {
	return "address"
}
